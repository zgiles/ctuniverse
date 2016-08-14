protocol..

# Overall concepts

Several message types will be present, each will be wrapped in a SpaceMessage.
The `messagetype` field will be a string that states the type of the object that is being wrapped
The `o` field will be the object itself. As this is JSON, it will be a dictionary (object) with keys / values as-per its type.

# SpaceMessage
SpaceMessage =
{
 messagetype: ("SpaceObject" | "SpaceControl" | "SpaceID"),
 o: (SpaceObject | SpaceControl | SpaceID)
}


# SpaceObject
The SpaceObject is based on the space game's SpaceObject so they can be directly ported between.
Today that makes this object as follows:

SpaceObject =
{
	uuid: String,
  owner: String,
  type: String,
  global: [ Float64 ... ],
  velocity: [ Float64 ... ],
  angle: Float64,
  angle_velocity: Float64,
  boost: Int64,
  thrusters: [ { type: String, firing: Int64 } ... ]
}

# SpaceControl

# SpaceID
