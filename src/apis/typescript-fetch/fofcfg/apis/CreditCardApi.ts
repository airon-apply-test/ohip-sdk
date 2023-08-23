/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  CreditCardInfo,
  ExceptionDetailType,
} from '../models';
import {
    CreditCardInfoFromJSON,
    CreditCardInfoToJSON,
    ExceptionDetailTypeFromJSON,
    ExceptionDetailTypeToJSON,
} from '../models';

export interface GetCreditCardInfoRequest {
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    reservationId?: string;
    idExtension?: number;
    idContext?: string;
    type?: string;
    cardId?: string;
    cardIdExtension?: number;
    cardIdContext?: string;
    cardIdType?: string;
    accessTransactionType?: GetCreditCardInfoAccessTransactionTypeEnum;
    xExternalsystem?: string;
    acceptLanguage?: string;
}

/**
 * 
 */
export class CreditCardApi extends runtime.BaseAPI {

    /**
     * Use this API to return the token inclusive cardType and expirationDate for the credit card used in a specific reservation. Oracle does not provide any credit card numbers. Include the hotelId and the cardId in the parameters. The cardId can be returned with the getReservation operation. <p><strong>OperationId:</strong>getCreditCardInfo</p>
     * Get credit card token information
     */
    async getCreditCardInfoRaw(requestParameters: GetCreditCardInfoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreditCardInfo>> {
        const queryParameters: any = {};

        if (requestParameters.hotelId !== undefined) {
            queryParameters['hotelId'] = requestParameters.hotelId;
        }

        if (requestParameters.reservationId !== undefined) {
            queryParameters['reservationId'] = requestParameters.reservationId;
        }

        if (requestParameters.idExtension !== undefined) {
            queryParameters['idExtension'] = requestParameters.idExtension;
        }

        if (requestParameters.idContext !== undefined) {
            queryParameters['idContext'] = requestParameters.idContext;
        }

        if (requestParameters.type !== undefined) {
            queryParameters['type'] = requestParameters.type;
        }

        if (requestParameters.cardId !== undefined) {
            queryParameters['cardId'] = requestParameters.cardId;
        }

        if (requestParameters.cardIdExtension !== undefined) {
            queryParameters['cardIdExtension'] = requestParameters.cardIdExtension;
        }

        if (requestParameters.cardIdContext !== undefined) {
            queryParameters['cardIdContext'] = requestParameters.cardIdContext;
        }

        if (requestParameters.cardIdType !== undefined) {
            queryParameters['cardIdType'] = requestParameters.cardIdType;
        }

        if (requestParameters.accessTransactionType !== undefined) {
            queryParameters['accessTransactionType'] = requestParameters.accessTransactionType;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }

        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }

        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }

        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }

        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }

        const response = await this.request({
            path: `/creditCardInfo`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreditCardInfoFromJSON(jsonValue));
    }

    /**
     * Use this API to return the token inclusive cardType and expirationDate for the credit card used in a specific reservation. Oracle does not provide any credit card numbers. Include the hotelId and the cardId in the parameters. The cardId can be returned with the getReservation operation. <p><strong>OperationId:</strong>getCreditCardInfo</p>
     * Get credit card token information
     */
    async getCreditCardInfo(requestParameters: GetCreditCardInfoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreditCardInfo> {
        const response = await this.getCreditCardInfoRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const GetCreditCardInfoAccessTransactionTypeEnum = {
    Additional: 'Additional',
    ArPayment: 'ArPayment',
    Batch: 'Batch',
    BatchDeposit: 'BatchDeposit',
    BatchDepositRefund: 'BatchDepositRefund',
    Billing: 'Billing',
    CheckIn: 'CheckIn',
    CheckInManualAuth: 'CheckInManualAuth',
    Deposit: 'Deposit',
    EndOfDay: 'EndOfDay',
    EndOfDayDeposit: 'EndOfDayDeposit',
    Manual: 'Manual',
    Other: 'Other',
    PasserBy: 'PasserBy',
    PostIt: 'PostIt',
    PreCheckIn: 'PreCheckIn',
    Refund: 'Refund',
    Scheduled: 'Scheduled'
} as const;
export type GetCreditCardInfoAccessTransactionTypeEnum = typeof GetCreditCardInfoAccessTransactionTypeEnum[keyof typeof GetCreditCardInfoAccessTransactionTypeEnum];
