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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 40.0
    z: -0.1
  }
  scale {
    x: -0.111584
    y: 0.050779
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 31.0
    z: -0.1
  }
  scale {
    x: 0.12425
    y: 0.051014
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 45.0
    z: -0.1
  }
  scale {
    x: 0.106381
    y: 0.050779
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 53.0
    z: -0.1
  }
  scale {
    x: -0.118457
    y: 0.051014
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 64.0
    z: -0.1
  }
  scale {
    x: 0.127726
    y: 0.051014
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 23.0
    z: -0.1
  }
  scale {
    x: -0.117336
    y: 0.051014
    z: 0.96
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 42.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 31.003529\n"
  "  data: 6.104739\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
