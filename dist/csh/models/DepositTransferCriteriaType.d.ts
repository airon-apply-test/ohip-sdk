/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DepositTransferToReservationType } from './DepositTransferToReservationType';
import type { ReservationId } from './ReservationId';
/**
 * Criteria type for transfer of deposit amount from one reservation to one or more reservation(s).
 * @export
 * @interface DepositTransferCriteriaType
 */
export interface DepositTransferCriteriaType {
    /**
     * Property where the reservation exists.
     * @type {string}
     * @memberof DepositTransferCriteriaType
     */
    hotelId?: string;
    /**
     *
     * @type {ReservationId}
     * @memberof DepositTransferCriteriaType
     */
    reservationId?: ReservationId;
    /**
     * The reservation id and the amount to transfer detail.
     * @type {Array<DepositTransferToReservationType>}
     * @memberof DepositTransferCriteriaType
     */
    toReservations?: Array<DepositTransferToReservationType>;
    /**
     * User Comments for the transfer operation.
     * @type {string}
     * @memberof DepositTransferCriteriaType
     */
    comments?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof DepositTransferCriteriaType
     */
    cashierId?: number;
}
/**
 * Check if a given object implements the DepositTransferCriteriaType interface.
 */
export declare function instanceOfDepositTransferCriteriaType(value: object): boolean;
export declare function DepositTransferCriteriaTypeFromJSON(json: any): DepositTransferCriteriaType;
export declare function DepositTransferCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DepositTransferCriteriaType;
export declare function DepositTransferCriteriaTypeToJSON(value?: DepositTransferCriteriaType | null): any;
