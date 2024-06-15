import { HttpClient, HttpHeaders, HttpRequest, HttpResponse, HttpParams, HttpEvent } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { catchError, filter, map } from 'rxjs/operators';

export enum RequestMethod {
  Get = 'GET',
  Head = 'HEAD',
  Post = 'POST',
  Put = 'PUT',
  Delete = 'DELETE',
  Options = 'OPTIONS',
  Patch = 'PATCH'

}

@Injectable()
export class ApiService {

  headers = new HttpHeaders({
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  });

  constructor(private http: HttpClient) { }

  get(path: string, args?: any): Observable<any> {
    let options = {
      headers: this.headers,
      params: new HttpParams()
    };

    if (args) {
      options['params'] = this.serialize(args);
    }

    return this.http.get(path, options)
      .pipe(catchError(this.checkError.bind(this)));
  }

  post(url: string, data: any, customHeaders?: HttpHeaders): Observable<any> {
    let headers = new HttpHeaders();
    if (!(data instanceof FormData)) {
      headers = headers.set('Content-Type', 'application/json');
    }
    
    if (customHeaders) {
      customHeaders.keys().forEach(key => {
        headers = headers.set(key, customHeaders.get(key) || '');
      });
    }

    return this.http.post(url, data, { headers });
  }

  put(path: string, body?: any): Observable<any> {
    return this.request(path, body, RequestMethod.Put);
  }
 
  delete(path: string, body?: any): Observable<any> {
    return this.request(path, body, RequestMethod.Delete);
  }
  patch(path: string, body?: any): Observable<any> {
    return this.request(path, body, RequestMethod.Patch);
  }
  
  
  private request(path: string, body: any, method = RequestMethod.Post, customHeaders?: HttpHeaders): Observable<any> {
    const req = new HttpRequest(method, path, body, {
      headers: customHeaders || this.headers,
      responseType: 'json'
    });

    return this.http.request(req).pipe(
      filter((response: HttpEvent<any>): response is HttpResponse<any> => response instanceof HttpResponse),
      map((response: HttpResponse<any>) => response.body),
      catchError(error => this.checkError(error))
    );
  }

  private checkError(error: any): any {
    throw error;
  }

  private serialize(obj: any): HttpParams {
    let params = new HttpParams();

    for (const key in obj) {
      if (obj.hasOwnProperty(key) && !this.looseInvalid(obj[key])) {
        params = params.set(key, obj[key]);
      }
    }

    return params;
  }

  private looseInvalid(a: string | number): boolean {
    return a === '' || a === null || a === undefined;
  }
}