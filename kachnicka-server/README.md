API server to Kachnicka MKRFox1200

	gcloud components update
	$(dirname -- "$(readlink $(which gcloud))")/dev_appserver.py app.yaml
	gcloud app deploy

## Installation of Google Cloud SDK

### Mac OSX

	brew cask info google-cloud-sdk
	brew tap caskroom/cask
	brew cask install google-cloud-sdk