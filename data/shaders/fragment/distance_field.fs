uniform sampler2D tex;
uniform float dist_min;
uniform float dist_max;

void main() {
  vec2 tpos = gl_TexCoord[0].xy;
  float dist = texture2D(tex, tpos).a;
  float x = 0.0;
  if (dist > 0.1 && dist <= 0.2) {
    x = 1.0;
  }
  if (dist > 0.3 && dist <= 0.4) {
    x = 1.0;
  }
  if (dist > 0.5 && dist <= 0.6) {
    x = 1.0;
  }
  if (dist > 0.7 && dist <= 0.8) {
    x = 1.0;
  }
  gl_FragColor = gl_Color * vec4(1.0, 1.0, 1.0, x);
  return;
  float alpha = smoothstep(dist_min, dist_max, dist);
  gl_FragColor = gl_Color * vec4(1.0, 1.0, 1.0, alpha);
}
