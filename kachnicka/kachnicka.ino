#include <ArduinoLowPower.h>
#include <SigFox.h>

// We want to send a boolean value to signal a binary event
// like open/close or on/off

bool value_to_send = true;
int sleepMins = 20;

#define DEBUG 0

void setup() {

  if (DEBUG) {
    Serial.begin(9600);
    while (!Serial) {};
  }

  //  if (!SigFox.begin()) {
  //    //something is really wrong, try rebooting
  //    reboot();
  //  }

  //  SigFox.begin();
  //  SigFox.debug();

  if (DEBUG) {
    Serial.print("Setup done");
  }
}

void loop() {

  sendString();

  if (DEBUG) {
    Serial.print("Sleeping now minutes: ");
    Serial.println(sleepMins);
  }
  // LowPower.sleep(sleepMins * 60 * 1000);

  // sleep XXX mins
  delay(sleepMins * 60 * 1000);
}

void sendString() {
  // Start the module
  SigFox.begin();
  SigFox.debug();
  // Wait at least 30mS after first configuration (100mS before)
  delay(100);
  // Clears all pending interrupts
  SigFox.status();
  delay(100);

  int temp = SigFox.internalTemperature();
  SigFox.beginPacket();
  SigFox.print(temp);
  int ret = SigFox.endPacket();  // send buffer to SIGFOX network
  SigFox.end();

  if (DEBUG) {
    if (ret > 0) {
      Serial.println("No transmission");
    } else {
      Serial.println("Transmission ok");
      Serial.println(ret);
      Serial.println(temp);
    }
  }

}

//void reboot() {
//  NVIC_SystemReset();
//  while (1);
//}


