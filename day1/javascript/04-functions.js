function nothing() {
  console.log("nothing");
}

nothing();

function greet(name) {
  console.log(`Hello, ${name}`);
}

greet("Jaber");

function square(n) {
  return n * n;
}

console.log(square(4));

const nothing2 = () => {
  console.log("nothing2");
};

nothing2();

const add = (x, y, z) => {
  return x + y + z;
};

console.log(add(1, 2, 3));
