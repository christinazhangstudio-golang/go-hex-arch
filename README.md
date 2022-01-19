Hexagonal Architecture is pattern for designing an application that decouples application components. For example, an HTTP component can be plugged in, or GRPC, to drive the application. The application does not need to depend on the components that drive it, or the components that it drives of makes calls to. Commmunication between these components is achieved through ports and adapters. Adapters can drive (e.g. HTTP receives calls, causing app to do something) or be driven by application (e.g. queries DB). Also called primary/secondary adapters. There's three layers to app: (1) Domain layer: contains business/domain logic (2) Application layer: orchestrate use of domain code, adapt requests from framework to domain layer (3) Framework layer: logic for outside components (e.g. database, GRPC). These layers are also decoupled; dependencies point inward (layers depend on inside layer, and not other direction e.g. domain layer cannot depend on application). -> To communicate with database from application layer, need to use dependency injection - means: instead of calling framework layer to call an instance of database from application code, just invert control of database instantiation to program on start-up, and have the program inject the database instantiation in application code on start up.

Structure (high-level):
    - cmd
        - main.go (for start-up of program, does dependency injection)
    - internal
        - adapters
            - app (application layer)
            - core (domain layer)
            - framework (framework layer)
                - left (GRPC and HTTP code)
                - right (database code)
        - ports (contain ports organized by layers, left/right framework ports will be added)

ports define interface
    - core.go defines ArithmeticPort, implemented by the struct Adapter in arithmetic.go
    - app.go defines APIPort, implemented by the struct Adapter in app.go
    - framework_right.go defines DbPort, implemented by the struct Adapter in db.go

main creates an instance for a port (of ArithmeticPort type), and assigns it an Arithmetic adapeter, using the NewAdapter() method in arithmetic.go

api.go has an attribute of ports.ArithmeticPort type, which allows it to do dependency injection, allowing it to access methods from the ArithmeticPort interface

If there are any changes to core layer, don't need to worry about application layer; core layer (arithmetic.go) implements methods in core.go - application layer and core layer decoupled.

Application layer needs access to database adapter, so in api.go, db ports.DbPort is an attribute in Adapter struct - whenever a new Adapter is created for the API, we can inject db into it

The application layer will be able to interact with both the core and framework layer, but since dependencies point inward, have to give application layer access to database through dependency injection (also giving application layer access to core through dependency injection).

-----
GRPC:

RPC (Remote Procedure Call) - allows computer to send/return responses (execute procedures) on another computer/address space
GRPC uses HTTP/2 protocol - binary instead of textual like HTTP/1, making transfer/parsing data more machine-friendly (faster, more efficient)

GRPC uses protocol buffers, which structure serialized data. Proto bufs involve using a language to describe structure of data and a program that generates source code from that description - that source code is used to write/read that serialized data.

-----
func main is the start of the application, and within the function, the dependencies are initiated. Necessary dependencies are injected in the adapters that need it (e.g. appAdapter has dbaseAdapter and core injected into it / gRPCAdapter has appAdapter injected into it).

Dependency injection is so that we can have dependencies point inwards.