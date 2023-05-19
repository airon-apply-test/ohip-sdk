# OHIPS SDK

### Usage

```typescript
import { Api } from 'ohips-sdk';

const api = new Api({
  hostName: '<Host Name>',
  appKey: '<App Key>',
  clientId: '<Client ID>',
  clientSecret: '<Client Secret>',
  username: '<Username>',
  password: '<Password>',
});

const response = await api.client.rtp.packages.getPackages('<HotelId>', {
  includeGroup: true,
  sellSeparate: true,
  startDate: '2021-09-26',
  endDate: '2021-09-27',
  children: 0,
  adults: 1,
  fetchInstructions: ['Header', 'CalculatedPrice', 'Items'],
  hotelId: ['<HotelId>'],
  ticketPostingRhythm: false,
});

console.log(response);
```

### Development

### Requirements

- Close [oracle/hospitality-api-docs](https://github.com/oracle/hospitality-api-docs) in the projects root directory.
- Change branch to `property_22.3`

### Building

- run `yarn build`
