import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {ListComponent} from './components/resources/list/list.component';
import {MatTableModule} from '@angular/material/table';
import {StatusComponent} from './components/status/status.component';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatButtonModule} from '@angular/material/button';
import {HttpClientModule} from '@angular/common/http';
import {ControlButtonsComponent} from './components/api/control-buttons/control-buttons.component';
import {MatSortModule} from '@angular/material/sort';
import {StoreModule} from '@ngrx/store';
import {SearchDrawerComponent} from './components/resources/search-drawer/search-drawer.component';
import {NgrxFormsModule} from 'ngrx-forms';
import {metaReducers, reducers} from './reducers';
import {StoreDevtoolsModule} from '@ngrx/store-devtools';
import {environment} from '../environments/environment';
import {ResourceCountComponent} from './components/resources/resource-count/resource-count.component';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipsModule} from '@angular/material/chips';
import {MatIconModule} from '@angular/material/icon';

@NgModule({
  declarations: [
    AppComponent,
    ListComponent,
    StatusComponent,
    ControlButtonsComponent,
    SearchDrawerComponent,
    ResourceCountComponent
  ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        HttpClientModule,

        MatTableModule,
        MatPaginatorModule,
        MatProgressSpinnerModule,
        MatFormFieldModule,
        MatInputModule,
        MatToolbarModule,
        MatSidenavModule,
        MatButtonModule,
        MatSortModule,
        NgrxFormsModule,

        // ngrx
        StoreModule.forRoot(reducers, {metaReducers}),
        !environment.production ? StoreDevtoolsModule.instrument() : [],
        StoreDevtoolsModule.instrument({maxAge: 25, logOnly: environment.production}),
        MatAutocompleteModule,
        MatChipsModule,
        MatIconModule,
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
