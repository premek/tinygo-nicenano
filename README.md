# tinygo on nice!nano board example

- Double-tap RST to GND pin
- The board should appear as a disk in /media/[USERNAME]/NICENANO
- `tinygo build -o /media/*/NICENANO/flash.uf2 -target feather-nrf52840 blinky1.go`
- The board should restart and the blue LED start blinking

It seems to build with feather-nrf52840 target for now.
