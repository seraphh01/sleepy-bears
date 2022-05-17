import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { LoginFormComponent } from './login-form/login-form.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { UserPageComponent } from './user-page/user-page.component';
import { AdminPageComponent } from './admin-page/admin-page.component';
import { UserProfileComponent } from './user-page/user-profile/user-profile.component';
import { StudentPageComponent } from './student-page/student-page.component';
import { TeacherPageComponent } from './teacher-page/teacher-page.component';
import { ChiefPageComponent } from './chief-page/chief-page.component';
import { RegisterUsersComponent } from './admin-page/register-users/register-users.component';
import { ViewUsersComponent } from './admin-page/view-users/view-users.component';
import { TokenInterceptor } from './Services/intercepter';
import { ProposeCourseComponent } from './teacher-page/propose-course/propose-course.component';
import {MatInputModule} from '@angular/material/input';
import { ViewProposedComponent } from './teacher-page/view-proposed/view-proposed.component';
import { ViewCoursesComponent } from './student-page/view-courses/view-courses.component';
import { EnrollOptionalComponent } from './student-page/enroll-optional/enroll-optional.component';
import { DeleteUserComponent } from './admin-page/delete-user/delete-user.component';
@NgModule({
  declarations: [
    AppComponent,
    LoginFormComponent,
    UserPageComponent,
    AdminPageComponent,
    UserProfileComponent,
    StudentPageComponent,
    TeacherPageComponent,
    ChiefPageComponent,
    RegisterUsersComponent,
    ViewUsersComponent,
    ProposeCourseComponent,
    ViewProposedComponent,
    ViewCoursesComponent,
    EnrollOptionalComponent,
    DeleteUserComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    AppRoutingModule,
    MatInputModule
  ],
  providers: [    {
    provide: HTTP_INTERCEPTORS,
    useClass: TokenInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
