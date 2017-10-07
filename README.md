
# Sigfox temperature sensor

Temperature sensor based on Arduino MKRFOX1200 https://www.arduino.cc/en/Main.ArduinoBoardMKRFox1200 communicating with cloud
over [Sigfox](www.gigfox.com) network with excellent [coverage](https://www.sigfox.com/en/coverage)

Sensor sends data to Google cloud appengine app linked using Sigfox portal's callback.

* `kachnicka` - Arduino's code for MKRFox1200 board with temp sensor. Schema TBD
* `kachnicka-server` - Google Appengine code in go for cloud
* `kachnicka-hacker` - load scripts for cloud to generate test data

Box designs for 3D printing:

* https://www.tinkercad.com/things/3PNvBc9iD1b-kachnicka-otocena
* https://www.tinkercad.com/things/2WrZpbAwtAm-kachnicka
