import { UserRole } from "./userRole";

export class User {
  id: string;
  username: string;
  password: string;
  email: string;
  name: string;
  lastname: string;
  userRole: UserRole;
  

  constructor(id: string, username: string, password: string, email:string, name: string, lastname: string, userRole: UserRole) {
    this.id = id;
    this.username = username;
    this.password = password;
    this.email = email;
    this.name = name;
    this.lastname = lastname;
    this.userRole = userRole; 
  }
}