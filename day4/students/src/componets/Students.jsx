import { useState } from "react";
import Student from "../classes/Student";

function Students({ students, setStudents }) {
  const [id, setId] = useState("");
  const [name, setName] = useState("");
  const [gpa, setGpa] = useState(0.0);

  const idOnchangeHandler = (e) => {
    setId(e.target.value);
  };

  const nameOnchangeHandler = (e) => {
    setName(e.target.value);
  };

  const gpaOnchangeHandler = (e) => {
    setGpa(parseFloat(e.target.value));
  };

  const addStudent = () => {
    students.push(new Student(id, name, gpa));
    setStudents(students.flat());
  };

  const removeStudent = (id) => {
    const newArr = students.filter((el) => {
      return el.id !== id;
    });

    // const newArr = [];
    // for (let i = 0; i < students.length; i++) {
    //     if (students[i].id !== id) {
    //         newArr.push(students[i]);
    //     }
    // }

    setStudents(newArr);
  };

  return (
    <div>
      <table>
        <thead>
          <tr>
            <td>ID</td>
            <td>Name</td>
            <td>GPA</td>
          </tr>
        </thead>

        <tbody>
          {students.map((student) => {
            return (
              <tr key={student.id}>
                <td>{student.id}</td>
                <td>{student.name}</td>
                <td>{student.gpa}</td>
                <td>
                  <button
                    onClick={() => {
                      removeStudent(student.id);
                    }}
                  >
                    🗑
                  </button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>

      <div>
        <label htmlFor="id">ID</label>
        <input id="id" value={id} onChange={idOnchangeHandler} />
        <br />
        <label htmlFor="name">Name</label>
        <input id="name" value={name} onChange={nameOnchangeHandler} />
        <br />
        <label htmlFor="gpa">GPA</label>
        <input
          id="gpa"
          type="number"
          value={gpa}
          onChange={gpaOnchangeHandler}
        />
        <br />
        <button onClick={addStudent}>Add Student</button>
      </div>
    </div>
  );
}

export default Students;
