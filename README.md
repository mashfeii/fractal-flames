# Fractal Flames

Simple fractal flames generator in Go.

## Usage

Project contains Makefile to build and run the project.

- `make build` - build the project
- `make bench` - run the benchmark (single vs multithreaded)
- `make clean` - clean results' directory

All the images are saved in the `results` directory. Program can be run multiple times at a time to generate several images.

### Usage flags

- `-w` - width of the image
- `-h` - height of the image
- `-tr` - list of transitions separated by comma ("1,2,3")
- `-n` - number of random vectors with coefficients and colors
- `-i` - number of iterations per sample
- `-s` - number of samples to generate picture
- `-g` - gamma correction coefficient
- `-sym` - number of symmetry axis
- `-t` - number of threads
- `-f` - format (png/jpeg)
- `-re` - static red channel value
- `-gr` - static green channel value
- `-bl` - static blue channel value

Example: `make build && ./bin/fractal_flame -n 30 -w 1000 -h 1000 -t 2 -tr 6,12,13`
