# Drawing a UML diagram for the MapReduce implementation in Go.

from plantuml import PlantUML
import tempfile

# Define UML Diagram
uml_code = """
@startuml
skinparam classAttributeIconSize 0

class "Main" {
    +main()
}

class "Map" {
    +Map(word: string, ch: channel)
}

class "Reduce" {
    +Reduce(frequencies: []map[string]int): map[string]int
}

Main --> Map : calls
Main --> Reduce : calls
Reduce <-- Map : sends data

@enduml
"""

# Create temporary file to store the UML diagram
with tempfile.NamedTemporaryFile(delete=False, suffix=".png") as temp_file:
    temp_file_name = temp_file.name

# Render the UML diagram
plantuml = PlantUML(url='http://www.plantuml.com/plantuml/img/')
plantuml.processes(uml_code, output=temp_file_name)

temp_file_name

