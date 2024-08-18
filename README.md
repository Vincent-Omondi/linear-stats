<a name="readme-top"></a>


# Linear Stats


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## About The Project
`linear-stats` is a command-line program designed to perform basic statistical calculations on a dataset. Specifically, the program calculates the Linear Regression Line and the Pearson Correlation Coefficient for a given set of data points. It reads the data from a file, computes the required statistics, and then prints the results in a specified format.

### Built With
<img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="60" height="60">



## Getting Started
To get started with the `linear-stats` program, follow the instructions below.

### Prerequisites
- Go installed on your local machine.
- A terminal or command prompt to run the script.


### Installation
1. Clone the repository:
```sh
git clone https://github.com/Vincent-Omondi/linear-stats.git
cd linear-stats
``` 

2. Build the Program:
If you want to build the program, you can do so by running:

```sh
go build -o linear-stats
```
3. Prepare Your Data:
Ensure that you have a file named `data.txt` with the data points you wish to analyze. The file should contain one numeric value per line.

## Usage

Once you have the necessary files and Go installed, you can run the program using the following command:

```sh
go run main.go data.txt
```
Or, if you built the program:

```sh
./linear-stats data.txt
```

This command will output:

- The Linear Regression Line in the form y = <value>x + <value>.
- The Pearson Correlation Coefficient with 10 decimal places of precision.

Example Output
```sh
Linear Regression Line: y = 2.345678x + 1.234567
Pearson Correlation Coefficient: 0.9876543210
```


## RoadMap
 - Add support for different types of statistical calculations.
 - Allow the user to specify the output file for results.
 - Implement unit tests to ensure the accuracy of the calculations.
 - Provide a GUI interface for users who prefer not to use the command line.


## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.


  1. Fork the Project
  2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
  3. Commit your Changes (git commit -m 'Add some AmazingFeature')
  4.  Push to the Branch (git push origin feature/AmazingFeature)
  5. Open a Pull Request


## License
This project is licensed under the MIT License.


## Contact
If you have any questions or need further assistance, please feel free to contact the project maintainer:

[X](https://twitter.com/vinomondi_1/)

[Github](https://github.com/Vincent-Omondi/)


## Acknowledgments

Special thanks to the Zone01 for their valuable resources and support.

<p align="right">(<a href="#linear-stats">back to top</a>)</p>
