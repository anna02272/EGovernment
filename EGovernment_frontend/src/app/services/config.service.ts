import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class ConfigService {
  private _vehicles_api_url = 'http://localhost:8000/api';
  private _statistics_api_url = 'http://localhost:8082/api';

  private _vehicles_url = this._vehicles_api_url + '/vehicles';
  private _statistics_url = this._statistics_api_url + '/statistics';

  get vehicles_url(): string {
    return this._vehicles_url;
  }

  get statistics_url() {
    return this._statistics_url;
  }
}

