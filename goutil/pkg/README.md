# pkg

## Flow

### memo

* Frameworks and Drivers
  * UI: CLI(Standard Output)
* Interface Adapters
  * Data Input/Persistence/Output
    * CRUD(Create/Read/Update/Delete)
    * Input: Data processing for Application Business Rules
    * Persistense: Save data
    * Output: Result (outputs to standard output)
* Application Business Rules
  * This layer indicates "What can software do?".
  * uses objects in Enterprise Business Rules for achieving usecase
* Enterprise Business Rules
  * Entity
  * Objects that expresses business rogic

### Example

* Frameworks and Drivers
  * UI
    * CLI (standard output)
  * External Interfaces `<Struct>`
    * http request
    * cloud api
    * kubernetes api
* Interface Adapters
  * Contorller `<Struct>`
    * User input data (e.g. cmd args) processing
  * Gateways(Repository) `<Interface>,<Struct>`
    * CRUD (DB, files)
      * configRepository `<struct>`
      * httpRepository `<struct>`
      * cloudapiRepositoy `<struct>`
      * k8sRepository `<struct>`
      * includes external interfaces
        * http request `<Interface>` (DIP)
        * cloud api `<Interface>` (DIP)
        * kubernetes api `<Interface>` (DIP)
  * Presenter `<Struct>`
    * Output to standard output
      * errorf
* Application Business Rules
  * UseCases
    * UseCase `<Interface>`
      * subcommands
    * Interactor `<Struct>`
      * struct implements usecase
      * includes repository and presenter interfaces (DIP)
        * httpRepository `<Interface>` (DIP)
        * cloudapiRepository `<Interface>` (DIP)
        * k8sRepository `<Interface>` (DIP)
* Enterprise Business Rules
  * Entity `<struct>`
    * config
    * postdata


### Overview

* UI -> main -> UserInput -> Controller
  * Controller edits input data
    * input data example: subcmd, userName
  * Controllerは複数種類を用意
    * DIcontainer
    * 利用するDB等で使い分け
* Controller -> UseCase(=interface) -> Interactor
  * class InteractorはRepositoryおよびPresenterのInterfaceを受け取り
    * RepositoryおよびPresenterは上位層なので、Interface(DIP)を利用することで依存ルールを順守
* Interactor -> Repository -> RepositoryImpl
  * DB access (CRUD)
* Interactor -> Presenter -> PresenterImpl
  * Output to stdout

## Frameworks and Drivers Layer

* DB
* Web
* Devices
* UI
  * cli

## Interface Adapters Layer

* Controllers
  * subcmd/
* Gateways (Repository)
  * CRUD
    * config/
* Presenters
  * 表示のためのデータ加工
  * log/

## Application Business Rules Layer

* Use Cases
  * ユースケース(Interface)
  * InteractorはUsecasesを実装
    * ビジネスロジックを表すのではなく、Entityの調整に注力(RepositoryやPresenterでControll)
	* RepositoryやPresenter(上Layerの情報)はInterfaceで受け取り

## Enterprise Business Rules:

* Entities
  * config