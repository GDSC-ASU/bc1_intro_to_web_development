export function exp2(power) {
  return 1 << power;
}

export default class Car {
  model;
  speed;
  constructor(model, speed) {
    this.model = model;
    this.speed = speed;
  }
}
