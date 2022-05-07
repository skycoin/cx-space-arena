# cx-space-arena 

## Controls

Arrow keys - control the spaceship
Space bar - Fire a laser
Key A - Spawn an asteroid

## Dependencies
You will need the latest version of Golang, instructions for installation are listed [here](https://go.dev/doc/install)

The 2D library used is [Ebiten](https://ebiten.org/), which has the following dependencies:

Debian / Ubuntu
```
sudo apt install libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
```
Fedora
```
sudo dnf install mesa-libGLU-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config
```
Solus
```
sudo eopkg install libglu-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel libxxf86vm-devel alsa-lib-devel pkg-config
```
Arch
```
sudo pacman -S mesa libxrandr libxcursor libxinerama libxi pkg-config
```
Alpine
```
sudo apk add alsa-lib-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev pkgconf
```

## Building and running

Once the dependencies are installed, clone the directory and cd into it:
```
git clone https://github.com/skycoin/cx-space-arena.git
cd cx-space-arena
```
Then, you can either run the game directly:
```
go run main.go
```
Or build it and run the executable:
```
go build
./cx-space-arena
```
