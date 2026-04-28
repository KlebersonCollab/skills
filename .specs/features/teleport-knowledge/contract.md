# Validation Contract: Teleport Knowledge

## Validation Sensors
- [x] **Sensor: MD Serialization**: Must correctly pack LEARNINGS/MEMORY into JSON.
- [x] **Sensor: Storage Integrity**: Global hub files must follow `feature_ID.json` format.
- [x] **Sensor: Merge Logic**: Import must append to `LEARNINGS.md` without overwriting existing content.
- [x] **Sensor: Context Load**: Must identify active feature path correctly from root.
