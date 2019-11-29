import {Component, OnInit} from '@angular/core';
import {createWorker, ImageLike} from 'tesseract.js';
import {getDocument, PDFSource} from 'pdfjs-dist/webpack';

@Component({
  selector: 'app-documents',
  templateUrl: './documents.component.html',
  styleUrls: ['./documents.component.scss']
})
export class DocumentsComponent implements OnInit {

  public files: Set<File> = new Set();
  text: string;

  constructor() {
  }

  ngOnInit() {
    // this.doOCR();
  }

  async doOCR() {
    const worker = createWorker({
      logger: m => console.log(m),
    });
    await worker.load();
    await worker.loadLanguage('eng');
    await worker.initialize('eng');
    const {data: {text}} = await worker.recognize('https://tesseract.projectnaptha.com/img/eng_bw.png');
    console.log(text);
    await worker.terminate();
  }

  async doOCRImage(image: ImageLike) {
    const worker = createWorker({
      logger: m => console.log(m),
    });
    await worker.load();
    await worker.loadLanguage('deu');
    await worker.initialize('deu');
    const {data: {text}} = await worker.recognize(image);
    console.log(text);
    this.text = text;
    await worker.terminate();
  }

  onFilesAdded(event: any) {
    console.log(event);

    if (event.target.files.length > 0) {
      const file = event.target.files[0] as File;
      console.log(file);
      console.log(file.name);
      console.log(file.type);

      const fileReader = new FileReader();

      fileReader.onload = () => {
        let foo = getDocument(fileReader.result as PDFSource);
        foo.promise.then((pdf) => {
          console.log(pdf);
          pdf.getPage(1).then((page) => {
            let viewport = page.getViewport({scale: 1.5});
            //
            // Prepare canvas using PDF page dimensions
            //

            let canvas = document.getElementById('pdf-to-image-canvas') as HTMLCanvasElement;
            let context = canvas.getContext('2d');
            canvas.height = viewport.height;
            canvas.width = viewport.width;
            //
            // Render PDF page into canvas context
            //
            let task = page.render({canvasContext: context, viewport: viewport});
            task.promise.then(() => {
              console.log(canvas.toDataURL('image/jpeg'));
              this.doOCRImage(canvas.toDataURL('image/jpeg'));
            });
          });
        });
      };
      fileReader.readAsArrayBuffer(file);

      // this.doOCRImage(file);
    }
  }
}
