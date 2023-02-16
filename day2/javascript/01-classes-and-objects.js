let obj = {
  name: "Jaber",
  phone: "+962791234567",
};

console.log("whole object", obj);

console.log("name", obj.name);

obj.name = "Ali";
console.log("name", obj.name);

class Student {
    id;
    name;
    gpa;
    courses;
    /** 
     * @param{string} id 
     * @param{string} name 
     * @param{number} gpa
     * @param{Array} courses
     */
    constructor (id, name, gpa, courses) {
        this.id = id;
        this.name = name;
        this.gpa = gpa;
        this.courses = courses;
    }
}
let ali = new Student("1", "Ali", 90, ["C++", "Discrete"]);
let jaber = new Student("2", "Jaber", 70, ["C++", "Discrete", "English"]);

console.log(ali);
console.log(jaber);
