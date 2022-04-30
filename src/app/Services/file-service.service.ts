import { Injectable } from '@angular/core';
import { delay, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileService {

  constructor() { }

  readFileContent(file: File): Promise<string> {
    return new Promise<string>((resolve, reject) => {
        if (!file) {
            resolve('');
        }

        const reader = new FileReader();

        reader.onload = (e) => {
            const text = reader.result!.toString();
            resolve(text);
        };

        reader.readAsText(file);
    });
}

  downloadTextFile(content: string, extension: string){
    const link = document.createElement('a');
    let blob = new Blob([content]);
    link.setAttribute('target', '_blank');
    link.setAttribute('href', URL.createObjectURL(blob));
    link.setAttribute('download', `file_name${extension}`);
    document.body.appendChild(link);
    link.click();
    link.remove();
  }
}
