export class Booking {
  id: string = '';
  title: string;
  date: string = new Date().toISOString();
  amount: number;
  categoryId: string;
  accountId: string;
  standingOrderId: string;
  standingOrderPeriod: string = 'MONTHLY';
  standingOrderLastDay: string;
  created: string;
  updated: string;
}
