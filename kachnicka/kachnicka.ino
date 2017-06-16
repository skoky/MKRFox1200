/*
  SigFox Send Boolean tutorial

  This sketch demonstrates how to send a simple binary data ( 0 or 1 ) using a MKRFox1200.
  If the application only needs to send one bit of information the transmission time
  (and thus power consumption) will be much lower than sending a full 12 bytes packet.

  This example code is in the public domain.
*/

#include <ArduinoLowPower.h>
#include <SigFox.h>

// We want to send a boolean value to signal a binary event
// like open/close or on/off

bool value_to_send = true;
int sleepMins = 10;

#define DEBUG 0

void setup() {

  if (DEBUG){
    Serial.begin(9600);
    while (!Serial) {};
  }

  // Initialize the SigFox module
  if (!SigFox.begin()) {
    if (DEBUG){
      Serial.println("Sigfox module unavailable !");
    }
    return;
  }

  // If we wanto to debug the application, print the device ID to easily find it in the backend
  if (DEBUG){
    SigFox.debug();
    Serial.println("Sigfox ready with ID  = " + SigFox.ID());
  }

}

void loop() {

  int temp = SigFox.internalTemperature();
  sendString(temp);
  
  Serial.print("Sleeping now minutes: ");
  Serial.println(sleepMins);
  LowPower.sleep(sleepMins * 60 * 1000);
  
}

void sendString(int str) {
  // Start the module
  SigFox.begin();
  // Wait at least 30mS after first configuration (100mS before)
  delay(100);
  // Clears all pending interrupts
  SigFox.status();
  delay(1);

  SigFox.beginPacket();
  SigFox.print(str);

  int ret = SigFox.endPacket();  // send buffer to SIGFOX network
  if (ret > 0) {
    Serial.println("No transmission");
  } else {
    Serial.println("Transmission ok");
  }

  Serial.println(SigFox.status(SIGFOX));
  Serial.println(SigFox.status(ATMEL));
  SigFox.end();
}

void sendStringAndGetResponse(int str) {
  // Start the module
  SigFox.begin();
  // Wait at least 30mS after first configuration (100mS before)
  delay(100);
  // Clears all pending interrupts
  SigFox.status();
  delay(1);

  SigFox.beginPacket();
  SigFox.print(str);

  int ret = SigFox.endPacket(true);  // send buffer to SIGFOX network and wait for a response
  if (ret > 0) {
    Serial.println("No transmission");
  } else {
    Serial.println("Transmission ok");
  }

  Serial.println(SigFox.status(SIGFOX));
  Serial.println(SigFox.status(ATMEL));

  if (SigFox.parsePacket()) {
    Serial.println("Response from server:");
    while (SigFox.available()) {
      Serial.print("0x");
      Serial.println(SigFox.read(), HEX);
    }
  } else {
    Serial.println("Could not get any response from the server");
    Serial.println("Check the SigFox coverage in your area");
    Serial.println("If you are indoor, check the 20dB coverage or move near a window");
  }
  Serial.println();

  SigFox.end();
}


