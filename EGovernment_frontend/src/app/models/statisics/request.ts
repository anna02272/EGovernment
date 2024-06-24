import {CategoryPerson} from "./categoryPerson";

export class Request {
  id: string;
  name: string;
  lastname: string;
  email: string;
  phone_number: number;
  category: CategoryPerson;
  question: string;
  date: string;

  constructor(id: string, name: string, lastname: string, email: string, phone_number: number, category: CategoryPerson, question: string, date: string) {
    this.id = id;
    this.name = name;
    this.lastname = lastname;
    this.email = email;
    this.phone_number = phone_number;
    this.category = category;
    this.question = question;
    this.date = date;
  }
}
