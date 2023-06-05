# Simulation of Solar Panel, Pump, and Storage Tank

This simulation aims to model the behavior of a solar panel, pump, and storage tank system. While the simulation is kept simple, it takes into account basic principles and concepts related to thermodynamics. It should be noted that this simulation may not cover all complexities and variations that exist in real-world systems.

## Solar Panel

The solar panel is represented by a simplified model with a few variables and functionalities. It captures solar energy based on the incident radiation and calculates the resulting temperature. The efficiency of the solar panel is considered to determine the amount of captured energy. The temperature of the solar panel is an important factor in heat transfer calculations.

## Pump

The pump component has limited functionality in this simulation. It serves to circulate the fluid within the system, facilitating heat transfer and maintaining the desired flow rate. 

## Storage Tank

The storage tank is assumed to be a water-based thermal energy storage system. It acts as a reservoir for storing thermal energy and has a certain volume. The temperature of the storage tank is a crucial parameter that affects the heat transfer process. The simulation considers heat transfer between the storage tank and other components.

Simulation Approach and Ambiguities

Given the limited knowledge in thermodynamics and the simplifications made, the simulation approach focuses on capturing the basic behavior of the system. It may not account for all real-world factors, complexities, or specific scenarios. Ambiguities in modeling and assumptions may exist due to the simplified nature of the simulation. Further research and expertise in thermodynamics would be necessary to develop a more accurate and comprehensive simulation.

Please note that this simulation serves as an introductory exploration and may not fully represent the intricacies of real-world solar panel, pump, and storage tank systems.

# Thermodynamics Basics

Thermodynamics is the branch of physics that deals with the relationships between heat, work, and energy. It provides a framework for understanding and analyzing the behavior of systems involving the transfer of heat and the performance of work.

## Key Concepts

### 1. Heat Transfer

Heat transfer is the process of thermal energy moving from one object or substance to another due to a temperature difference. There are three main methods of heat transfer:

- **Conduction:** Heat transfer through direct contact between molecules or particles.
- **Convection:** Heat transfer through the movement of a fluid (liquid or gas).
- **Radiation:** Heat transfer through electromagnetic waves.

### 2. Energy Conversion

Thermodynamics involves the conversion of energy from one form to another. The basic principles are:

- **First Law of Thermodynamics:** Energy cannot be created or destroyed; it can only be transferred or converted from one form to another.
- **Second Law of Thermodynamics:** The total entropy of an isolated system always increases over time.

### 3. System and Surroundings

In thermodynamics, a system refers to the portion of the universe under consideration, while the surroundings encompass everything outside the system. The system and surroundings interact through the transfer of heat and work.

## Heat Transfer Simulation Code

The provided code is a simulation of heat transfer involving a tank of water. It includes models and services to represent the tank and the pump responsible for water flow. Here's an overview of the code:

- `models` package: Contains the data models used in the simulation, such as `Storage Tank`, `Solar Panle` and `Pump`.
- `services` package: Provides services for performing operations on the system, which simulates the system of the solar panel, pump, and storage tank.
- `main` function: The entry point of the program where it called the method `StartApplication` where the simulation is executed.

# Code

## Simulation Overview

The simulation is implemented in the `application` function, where a `System` instance is created. The system comprises a solar panel, storage tank, and pump. It simulates the transfer of water based on solar radiation and other conditions.

## Solar Panel

The `SolarPanelService` interface defines the behavior of a solar panel, including capturing solar radiation, calculating energy conversion, and updating degradation. The `SolarPanel` struct implements this interface. It tracks the area, efficiency, temperature, degradation, and dust accumulation of the solar panel.

## Storage Tank

The `StorageTank` interface represents a water storage tank. The `StorageTank` struct implements this interface. It keeps track of the tank's volume, temperature, and water volume.

## Pump

The `Pump` interface defines the behavior of a water pump, such as pumping water from a storage tank. The `Pump` struct implements this interface. It has a flow rate that determines the speed of water transfer.

## Heat Transfer Simulation

The code simulates heat transfer within the system using the `Simulate` method of the `System` struct. Although the exact implementation of the heat transfer is not shown in the provided code snippet, the simulation likely involves calculating heat transfer coefficients, radiation, conduction, and convection effects based on temperature differences and other factors.

### Heat Transfer Function

The specific details of the heat transfer function are not included in the code snippet. However, it is an essential part of the simulation. The heat transfer function typically considers factors such as surface area, temperature differences, thermal conductivity, and convective heat transfer coefficients to calculate the amount of heat transferred between different components of the system.

## Running the Simulation

To run the simulation:

1. Ensure you have Go installed on your system.
2. Create a new Go project or navigate to an existing one.
3. Copy the code from github repo: [https://github.com/rudyjcruz831/heat-transfer-simulation.git]
4. Open a terminal or command prompt and navigate to the project directory.
5. Run the command `go run main.go` to execute the code.
6. To stop the continues simulation press (Ctrl+C)


# Unit Testing 

This package contains unit tests for the `services` module of the heat transfer simulation.

Please note that there is limited testing in this package due to the nature of the system being simulated. The tests provided serve as a sample and demonstrate the testing approach for specific methods.

## TestCaptureSolarEnergy

Verifies the `CaptureSolarEnergy` method in the `SolarPanelService` struct.

- **Test Function Signature**: `func TestCaptureSolarEnergy(t *testing.T)`
- **Test Scenario**:
  1. Create a new `SolarPanel` with predefined properties.
  2. Set the solar radiation and heat transfer coefficient.
  3. Create a new `SolarPanelService` with the created `SolarPanel`.
  4. Invoke `CaptureSolarEnergy` method on the `SolarPanelService`.
  5. Calculate the expected temperature based on the incident solar radiation.
  6. Assert that the actual temperature matches the expected temperature.

## TestUpdateDegradation

Verifies the `UpdateDegradation` method in the `SolarPanelService` struct.

- **Test Function Signature**: `func TestUpdateDegradation(t *testing.T)`
- **Test Scenario**:
  1. Create a new `SolarPanel` with predefined degradation.
  2. Create a new `SolarPanelService` with the created `SolarPanel`.
  3. Invoke `UpdateDegradation` method on the `SolarPanelService`.
  4. Calculate the expected degradation after the update.
  5. Assert that the actual degradation matches the expected degradation.

## TestTransferHeat

Verifies the `TransferHeat` method in the `StorageTankService` struct.

- **Test Function Signature**: `func TestTransferHeat(t *testing.T)`
- **Test Scenario**:
  1. Create a new `StorageTank` with predefined temperature and volume.
  2. Create a new `StorageTankService` with the created `StorageTank`.
  3. Create a new `SolarPanel` with predefined temperature and area.
  4. Set the heat transfer coefficient.
  5. Calculate the expected temperature after transferring heat.
  6. Invoke `TransferHeat` method on the `StorageTankService`.
  7. Assert that the actual temperature matches the expected temperature.

These tests serve as a basic demonstration of testing approaches for the heat transfer simulation services. Further testing can be done to cover additional scenarios and edge cases.

Please note that due to the nature of the simulation, there are limited aspects that can be thoroughly tested.
