let obj = {
  name: "Jaber",
  phone: "+962791234567",
};

console.log("whole object", obj);

console.log("name", obj.name);

obj.name = "Ali";
console.log("name", obj.name);

class Student {
  name;
  phone;
  gpa;
  constructor(name, phone, gpa) {
    this.name = name;
    this.phone = phone;
    this.gpa = gpa;
  }
}

let ali = new Student("Ali Al-Homsi", "+962791234567", 90.2);
console.log("ali's data", ali);
