import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class ConfigService {
  private _vehicles_api_url = 'http://localhost:8000/api';
  private _statistics_api_url = 'http://localhost:8082/api';
  private _auth_api_url = 'http://localhost:8085/api';

  private _vehicles_url = this._vehicles_api_url + '/vehicles';
  private _statistics_url = this._statistics_api_url + '/statistics';
  private _auth_url = this._auth_api_url + '/auth';
  private _user_url = this._auth_api_url + '/users';
  private _login_url = this._auth_url + '/login';
  private _register_url = this._auth_url + '/register';
  private _current_user_url = this._user_url + '/currentUser';

  get vehicles_url(): string {
    return this._vehicles_url;
  }

  get statistics_url() {
    return this._statistics_url;
  }

  get auth_url(): string {
    return this._auth_url;
  }

  get login_url(): string {
    return this._login_url;
  }

  get register_url(): string {
    return this._register_url;
  }

  get user_url(): string {
    return this._user_url;
  }

  get currentUser_url(): string {
    return this._current_user_url;
  }
}

