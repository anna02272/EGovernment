export class Response {
  id: string;
  text: string;
  attachment: string;
  accepted: boolean;
  send_to: string;
  date: string;
 
  constructor(id: string, text: string, attachment: string, accepted: boolean, send_to: string, date: string) {
    this.id = id;
    this.text = text;
    this.attachment = attachment;
    this.accepted = accepted;
    this.send_to = send_to;
    this.date = date;
  }
}