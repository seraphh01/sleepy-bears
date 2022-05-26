import { catchError, map, Observable, throwError } from "rxjs";

export  class Pipe {
    public static makePipe(func : Observable<any>): Observable<any> {
        return func.pipe(map(res => {
            return res;
          }), catchError(err => {
              return throwError(() => new Error(err.error.error));
            }
          ))
    }
}