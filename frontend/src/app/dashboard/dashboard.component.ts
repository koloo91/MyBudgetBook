import {Component, OnInit} from '@angular/core';
import {ErrorService} from '../services/error.service';
import {StatisticService} from '../services/statistic.service';
import {ErrorVo} from '../models/error.model';
import {ChartDataSets, ChartOptions, ChartType} from 'chart.js';
import {Label} from 'ng2-charts';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  isLoading: boolean = true;

  public barChartOptions: ChartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    // We use these empty structures as placeholders for dynamic theming.
    scales: {xAxes: [{}], yAxes: [{}]},
    plugins: {
      datalabels: {
        anchor: 'end',
        align: 'end',
      }
    }
  };
  public barChartLabels: Label[] = ['Dezember', 'November', 'Oktober', 'September', 'August', 'Juli', 'Juni', 'Mai', 'April', 'März', 'Februar', 'Januar'];
  public barChartType: ChartType = 'horizontalBar';
  public barChartLegend = true;

  public barChartData: ChartDataSets[] = [
    {data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], label: 'Ausgaben'},
    {data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], label: 'Einnahmen'}
  ];

  constructor(
    private statisticService: StatisticService,
    private errorService: ErrorService) {
  }

  ngOnInit() {
    this.loadData();
  }

  loadData() {
    this.statisticService.getMonthStatistics()
      .subscribe(data => {
          console.log(data);
          this.isLoading = false;
          this.barChartData[0].data = data.map(_ => Math.abs(_.expenses));
          this.barChartData[1].data = data.map(_ => _.incomes);
        },
        (error: ErrorVo) => {
          this.isLoading = false;
          this.errorService.showErrorMessage(error.message);
        });
  }
}
