#include <ArduinoLowPower.h>
#include <SigFox.h>

int sleepMins = 20;

#define DEBUG 0

void setup() {


  if (!SigFox.begin()) {
    Serial.println("Shield error or not present!");
    return;
  }

  if (DEBUG) {
    Serial.begin(9600);
    while (!Serial) {};
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
  delay(sleepMins * 60 * 1000);
  
  if (DEBUG) {
    Serial.println("Wake up");
  }
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
  SigFox.noDebug();
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


