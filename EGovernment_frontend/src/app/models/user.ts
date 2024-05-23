import { UserRole } from "./userRole";

export class User {
  id: string;
  username: string;
  password: string;
  email: string;
  jmbg: number;
  name: string;
  lastname: string;
  userRole: UserRole;
  

  constructor(id: string, username: string, password: string, email:string, jmbg: number, name: string, lastname: string, userRole: UserRole) {
    this.id = id;
    this.username = username;
    this.password = password;
    this.email = email;
    this.jmbg = jmbg;
    this.name = name;
    this.lastname = lastname;
    this.userRole = userRole; 
  }
}