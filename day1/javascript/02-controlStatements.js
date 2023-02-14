let shallIPass = true;
if (shallIPass) {
  console.log("You shall pass");
} else {
  console.log("You shall not pass");
}

let x = 12,
  y = 10,
  z = 13;

if (x > y && x > z) {
  console.log("x is the greatest");
} else if (y > x && y > z) {
  console.log("y is the greatest");
} else if (z > x && z > y) {
  console.log("z is the greatest");
}

let mark = "B";
switch (mark) {
  case "A":
    console.log("big nerd");
    break;
  case "B":
    console.log("nerd");
    break;
  case "C":
    console.log("ain't much");
    break;
  default:
    console.log("I don't know what is that!");
}

for (let i = 1; i <= 10; i++) {
  console.log(i);
}

let count = 10;
while (count > 0) {
  console.log(count);
  count--;
}
