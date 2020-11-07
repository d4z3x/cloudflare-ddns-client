package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/juju/errors"
	"github.com/mattolenik/cloudflare-ddns-client/ip"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	Token      string
	ZoneID     string
	Domain     string
	IP         string
	CF         *cloudflare.API
	TestBinary string
}

func TestIntegrationSuite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupSuite() {
	require := s.Require()
	var err error
	s.TestBinary = os.Getenv("TEST_BINARY")
	require.NotEmpty(s.TestBinary, "Integration tests require compiled binary of cloudflare-ddns at path specified by TEST_BINARY")

	s.Token = os.Getenv("CLOUDFLARE_TOKEN")
	require.NotEmpty(s.Token, "Integration tests require an API token specified by the CLOUDFLARE_TOKEN env var")

	s.Domain = os.Getenv("TEST_DOMAIN")
	require.NotEmpty(s.Domain, "Integration tests require a domain specified by the TEST_DOMAIN env var")

	s.IP, err = ip.GetPublicIP()
	require.NoError(err, "unable to get public IP for tests")

	s.CF, err = cloudflare.NewWithAPIToken(s.Token)
	if err != nil {
		require.NoError(err, "Unable to connect to CloudFlare, token may be invalid")
	}

	// Verify token before tests start
	_, err = s.CF.VerifyAPIToken()
	require.NoError(err, "CloudFlare token is not valid")

	s.ZoneID, err = s.CF.ZoneIDByName(s.Domain)
	require.NoErrorf(err, "Failed to get zone ID, error: %+v", err)
}

func (s *IntegrationSuite) TestWithConfigFile() {
	assert := s.Assert()
	require := s.Require()

	record := s.randomDNSRecord()
	defer s.deleteRecord(record)

	tmp := s.T().TempDir()
	configFile := path.Join(tmp, "config.toml")
	config := fmt.Sprintf(`
	domain="%s"
	record="%s"
	token="%s"`, s.Domain, record, s.Token)
	err := ioutil.WriteFile(configFile, []byte(config), 0644)
	assert.NoErrorf(err, "Could not write test config file")

	out, err := s.runProgram(nil, "--config", configFile)
	log.Println(out)
	require.NoError(err, "Expected no error when running cloudflare-ddns")
	assert.True(s.ipMatches(record), "Expected IP to have been updated")
}

func (s *IntegrationSuite) TestWithArguments() {
	assert := s.Assert()
	require := s.Require()

	record := s.randomDNSRecord()
	defer s.deleteRecord(record)

	out, err := s.runProgram(nil, "--domain", s.Domain, "--record", record, "--token", s.Token)
	log.Println(out)
	require.NoError(err, "Expected no error when running cloudflare-ddns")
	assert.True(s.ipMatches(record), "Expected IP to have been updated")
}

func (s *IntegrationSuite) TestExistingRecord() {
	assert := s.Assert()
	require := s.Require()

	record := s.createRandomDNSRecord("10.0.0.0")
	defer s.deleteRecord(record)

	out, err := s.runProgram(nil, "--domain", s.Domain, "--record", record, "--token", s.Token)
	log.Println(out)
	require.NoError(err, "Expected no error when running cloudflare-ddns")
	assert.True(s.ipMatches(record), "Expected IP to have been updated")
}

func (s *IntegrationSuite) TestWithEnvVars() {
	assert := s.Assert()
	require := s.Require()

	record := s.randomDNSRecord()
	defer s.deleteRecord(record)

	out, err := s.runProgram([]string{"DOMAIN=" + s.Domain, "RECORD=" + record, "TOKEN=" + s.Token})
	log.Println(out)
	require.NoError(err, "Expected no error when running cloudflare-ddns")
	assert.True(s.ipMatches(record), "Expected IP to have been updated")
}

func (s *IntegrationSuite) deleteRecord(record string) {
	r, err := s.getDNSRecordByName(record)
	s.Assert().NoErrorf(err, "Could not find DNS record of name '%s' in zone ID '%s', cannot clean up", record, s.ZoneID)
	err = s.CF.DeleteDNSRecord(s.ZoneID, r.ID)
	s.Assert().NoErrorf(err, "Failed to remove DNS record of name '%s' in zone ID '%s'", record, s.ZoneID)
}

func (s *IntegrationSuite) ipMatches(record string) bool {
	r, err := s.getDNSRecordByName(record)
	s.Require().NotNilf(r, "Expected record for '%s' to be not nil", record)
	s.Require().NoError(err, "Failed to get record ID of '%s'", record)
	return r.Content == s.IP
}

func (s *IntegrationSuite) getDNSRecordByName(record string) (*cloudflare.DNSRecord, error) {
	records, err := s.CF.DNSRecords(s.ZoneID, cloudflare.DNSRecord{Type: "A"})
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, r := range records {
		if r.Name == record {
			return &r, nil
		}
	}
	return nil, errors.NotFoundf("no record '%s' found in zone ID '%s'", record, s.ZoneID)
}

func (s *IntegrationSuite) createRandomDNSRecord(ip string) string {
	record := s.randomDNSRecord()
	_, err := s.CF.CreateDNSRecord(s.ZoneID, cloudflare.DNSRecord{
		Content: ip,
		Type:    "A",
		Name:    record,
	})
	s.Require().NoErrorf(err, "failed to create DNS record '%s' on domain '%s'", record, s.Domain)
	return record
}

func (s *IntegrationSuite) randomDNSRecord() string {
	return fmt.Sprintf("ddns-integration-test-%d.%s", rand.Intn(999999)+100000, s.Domain)
}

func (s *IntegrationSuite) runProgram(envVars []string, args ...string) (string, error) {
	cmd := exec.Command(s.TestBinary, args...)
	cmd.Env = envVars
	out, err := cmd.CombinedOutput()
	return string(out), err
}
