package engagebdr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

func TestEngageBDRSyncer(t *testing.T) {
	temp := template.Must(template.New("sync-template").Parse("https://match.bnmla.com/usersync/prbds2s?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}&r=localhost%2Fsetuid%3Fbidder%3Dgumgum%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D"))
	syncer := NewEngageBDRSyncer(temp)
	syncInfo, err := syncer.GetUsersyncInfo("1", "BOPVK28OVJoTBABABAENBs-AAAAhuAKAANAAoACwAGgAPAAxAB0AHgAQAAiABOADkA")
	assert.NoError(t, err)
	assert.Equal(t, "https://match.bnmla.com/usersync/prbds2s?gdpr=1&gdpr_consent=BOPVK28OVJoTBABABAENBs-AAAAhuAKAANAAoACwAGgAPAAxAB0AHgAQAAiABOADkA&r=localhost%2Fsetuid%3Fbidder%3Dgumgum%26gdpr%3D1%26gdpr_consent%3DBOPVK28OVJoTBABABAENBs-AAAAhuAKAANAAoACwAGgAPAAxAB0AHgAQAAiABOADkA%26uid%3D", syncInfo.URL)
	assert.Equal(t, "iframe", syncInfo.Type)
	assert.EqualValues(t, 62, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}