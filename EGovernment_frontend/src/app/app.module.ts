import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { HomeComponent } from './components/home/home.component';
import { ConfigService } from './services/config.service';
import { ApiService } from './services/api.service';
import { RefreshService } from './services/refresh.service';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
  // {
    // provide: HTTP_INTERCEPTORS,
    // useClass: TokenInterceptor,
  //   multi: true,
  // },
  ConfigService,
  ApiService,
  RefreshService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
