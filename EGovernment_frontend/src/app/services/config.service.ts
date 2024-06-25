import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class ConfigService {
  private _vehicles_api_url = 'http://localhost:8080/api/vehicle';
  private _statistics_api_url = 'http://localhost:8082/api';
  private _auth_api_url = 'http://localhost:8085/api';

  private _vehicles_url = this._vehicles_api_url + '/vehicles';
  private _auth_url = this._auth_api_url + '/auth';
  private _user_url = this._auth_api_url + '/users';

  private _login_url = this._auth_url + '/login';
  private _register_url = this._auth_url + '/register';
  private _current_user_url = this._user_url + '/currentUser';

  private _request_url = this._statistics_api_url + '/request';
  private _response_url = this._statistics_api_url + '/response';
  private _registeredVehiclesReport_url = this._statistics_api_url + '/registeredVehiclesReport';
  private _carAccidentDegreeReport_url = this._statistics_api_url + '/carAccidentDegreeReport';
  private _carAccidentTypeReport_url = this._statistics_api_url + '/carAccidentTypeReport';
  private _delictReport_url = this._statistics_api_url + '/delictReport';

  private _createVehicle_url = this._vehicles_api_url + '/createVehicle';
  private _allVehicles_url = this._vehicles_api_url + '/all';
  private _allRegisteredVehicles_url = this.allVehicles_url  + '/registered';
  private _getVehicleById = this._vehicles_api_url + '/get/';


  get vehicles_url(): string {
    return this._vehicles_url;
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

  get request_url(): string {
    return this._request_url;
  }

  get response_url(): string {
    return this._response_url;
  }

  get registeredVehiclesReport_url(): string {
    return this._registeredVehiclesReport_url;
  }

  get carAccidentDegreeReport_url(): string {
    return this._carAccidentDegreeReport_url;
  }

  get carAccidentTypeReport_url(): string {
    return this._carAccidentTypeReport_url;
  }

  get delictReport_url(): string {
    return this._delictReport_url;
  }

  get createVehicle_url(): string {
    return this._createVehicle_url;
  }

  get allVehicles_url(): string {
    return this._allVehicles_url;
  }

  get getVehicleById_url(): string {
    return this._getVehicleById;
  }

  get getRegisteredVehicles_url(): string {
    return this._allRegisteredVehicles_url;
  }

  getVehicleByCategoryAndYear_url(category: string, year: number): string {
    return `${this._vehicles_api_url}/get/category/${category}/year/${year}`;
  }
}

