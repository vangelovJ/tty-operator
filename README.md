<div id="top"></div>




<!-- PROJECT SHIELDS -->
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Tty-operator</h3>

  <p align="center">
    Kubernetes operator for running terminals inside the cluster
    <br />
</div>




<!-- ABOUT THE PROJECT -->
## About The Project


Here's why:
* I want to be able to create kubernete operators
* I am using gotty so that the terminal will work in the browser
* This project will further improve my coding skills


<p align="right">(<a href="#top">back to top</a>)</p>



### Built With


* [operator-sdk](https://sdk.operatorframework.io/)
* [Go](https://golang.org/)
* [Gotty](https://github.com/yudai/gotty)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* operator-sdk
* kubectl
* minikube
* go
* docker

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/vangelovJ/tty-operator.git
   ```
2. Build the docker image for the controller and push it to dockerhub
   ```sh
   make docker-build docker-push IMG="vangelovj/tty-operator:v0.0.3"
   ```
3. Deploy the operator locally (kubectl should be configured to access the k8s cluster)
   ```js
   make deploy
   ```

<p align="right">(<a href="#top">back to top</a>)</p>






<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Ven Angelov - vangelovj@jboxers.com

Project Link: [https://github.com/vangelovJ/tty-operator.git](https://github.com/vangelovJ/tty-operator.git)

<p align="right">(<a href="#top">back to top</a>)</p>





<p align="right">(<a href="#top">back to top</a>)</p>



