/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { PayeeInfoType } from './PayeeInfoType';
/**
 * Folio window view which holds the set of folios for a reservation.
 * @export
 * @interface ReservationFolioWindowType
 */
export interface ReservationFolioWindowType {
    /**
     *
     * @type {PayeeInfoType}
     * @memberof ReservationFolioWindowType
     */
    payeeInfo?: PayeeInfoType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ReservationFolioWindowType
     */
    balance?: CurrencyAmountType;
    /**
     * Payment Method Type
     * @type {string}
     * @memberof ReservationFolioWindowType
     */
    paymentMethod?: string;
    /**
     *
     * @type {number}
     * @memberof ReservationFolioWindowType
     */
    folioWindowNo?: number;
}
/**
 * Check if a given object implements the ReservationFolioWindowType interface.
 */
export declare function instanceOfReservationFolioWindowType(value: object): boolean;
export declare function ReservationFolioWindowTypeFromJSON(json: any): ReservationFolioWindowType;
export declare function ReservationFolioWindowTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationFolioWindowType;
export declare function ReservationFolioWindowTypeToJSON(value?: ReservationFolioWindowType | null): any;
