# Wallpaper Randomizer

Tool that will randomly pick a wallpaper in the provided directory.

*You can have it execute on system startup!*

**Dependencies**:
* swww
* go


### Configuring

```go
    // Simply change the placeholder variable below to your target directory.
  	wallpaperPath := "/home/user/Pictures/Wallpapers"
```

### Building

```
  git clone https://github.com/NM711/Wallpaper-Randomizer.git
```

```
  cd Wallpaper-Randomizer
```

```
  go build -o wallpaper_randomizer main.go
```

```
  # Move it to your local bin directory
  
  mv wallpaper_randomizer ~/.local/bin/
```

Make sure your "~/.local/bin" is read by your path.
