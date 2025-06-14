# GoAlpaca

Go implementation of Alpaca related devices and protocol


## Task List

- [ ] TUI implementation, perhaps also script interface (simple CLI).
- [ ] Simulator Devices (based on existing Alpaca Simulators).
- [ ] Persistance for DeviceTree; save restore server config and/or state.
- [ ] USB scan
- [ ] Drivers for:
  - [ ] Prima Luce Lab
    - [ ] Sesto Senso Focuser
    - [ ] Eagle Hardware (interface functions such as switches, labels, temperature and powerconsumption)
  - [ ] Celestion Mount (EDGE 8 Handcontroller)
  - [ ] Camera ZWO


## Reference Material

* https://www.ascom-standards.org/api/
* https://ascom-standards.org/AlpacaDeveloper/Index.htm
* https://ascom-standards.org/AlpacaDeveloper/ASCOMAlpacaAPIReference.html
* https://github.com/ASCOMInitiative/ASCOM.Alpaca.Simulators
* https://www.wireshark.org/#download
* https://gin-gonic.com/en/docs/
* https://go.dev/doc/tutorial/web-service-gin
* https://github.com/charmbracelet/bubbletea


Start on device interface for focuser, run against simulator, then reverse engineer a driver.

Tools for scanning usb, and UDP discover ... debug tools.
