components {
  id: "platform"
  component: "/level/platform.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 190.41339\n"
  "  data: 68.566246\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -1.0
    z: 1.0
  }
}
embedded_components {
  id: "spike"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -202.0
    z: -0.1
  }
}
embedded_components {
  id: "dangercollisions"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 193.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -185.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -105.0\n"
  "      y: -72.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: -0.7091444\n"
  "      w: 0.7050633\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 94.0\n"
  "      y: -72.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: -0.7091444\n"
  "      w: 0.7050633\n"
  "    }\n"
  "    index: 9\n"
  "    count: 3\n"
  "  }\n"
  "  data: 16.321613\n"
  "  data: 46.63574\n"
  "  data: 10.0\n"
  "  data: 16.321613\n"
  "  data: 46.63574\n"
  "  data: 10.0\n"
  "  data: 16.321613\n"
  "  data: 68.64499\n"
  "  data: 10.0\n"
  "  data: 16.321613\n"
  "  data: 68.64499\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "spike1"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 204.0
    z: -0.1
  }
  rotation {
    y: 0.96983147
    w: 0.24377629
  }
}
embedded_components {
  id: "spike2"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -104.0
    y: -85.0
    z: -0.1
  }
  rotation {
    z: 0.70090926
    w: 0.71325046
  }
}
embedded_components {
  id: "spike3"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 94.0
    y: -83.0
    z: -0.1
  }
  rotation {
    z: 0.70090926
    w: 0.71325046
  }
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/level/coin.go\"\n"
  ""
}
