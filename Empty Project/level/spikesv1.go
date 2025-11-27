components {
  id: "spike"
  component: "/level/spike.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 56.0
    z: -0.1
  }
  scale {
    x: 0.516302
    y: 0.336856
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 121.0
    z: -0.1
  }
  scale {
    x: -0.838633
    y: 0.325723
    z: 0.96
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 313.0
    z: -0.1
  }
  scale {
    x: 0.61831
    y: 0.336856
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 356.0
    z: -0.1
  }
  scale {
    x: -0.616805
    y: 0.336856
    z: 0.96
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 186.0
    z: -0.1
  }
  scale {
    x: 0.61831
    y: 0.336856
    z: 0.96
  }
}
embedded_components {
  id: "sprite5"
  type: "sprite"
  data: "default_animation: \"spike endlessrunner\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 350.0\n"
  "  y: 525.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 244.0
    z: -0.1
  }
  scale {
    x: 0.61831
    y: 0.336856
    z: 0.96
  }
}
embedded_components {
  id: "collisionobject"
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
  "      x: 202.0\n"
  "      y: -11.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 170.9768\n"
  "  data: 63.679367\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
