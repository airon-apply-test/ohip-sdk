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
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { PrepaidCardDetailsType } from './PrepaidCardDetailsType';
import type { ProfileId } from './ProfileId';
import type { ReservationId } from './ReservationId';
import type { SaleCriteriaType } from './SaleCriteriaType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Holds fixed charge information.
 * @export
 * @interface PrepaidCardCriteriaType
 */
export interface PrepaidCardCriteriaType {
    /**
     * Hotel context for the Reservations.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    hotelId?: string;
    /**
     *
     * @type {ReservationId}
     * @memberof PrepaidCardCriteriaType
     */
    reservationId?: ReservationId;
    /**
     * Family name, last name or Company Name.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    name?: string;
    /**
     * Given name, first name or names.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    firstName?: string;
    /**
     *
     * @type {ProfileId}
     * @memberof PrepaidCardCriteriaType
     */
    profileId?: ProfileId;
    /**
     * Prepaid card / account number.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    cardNo?: string;
    /**
     * Masked Prepaid card / account number.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    cardNumberMasked?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardCriteriaType
     */
    amount?: CurrencyAmountType;
    /**
     * Prepaid card pin code.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    pinCode?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof PrepaidCardCriteriaType
     */
    interfaceId?: UniqueIDType;
    /**
     *
     * @type {PrepaidCardDetailsType}
     * @memberof PrepaidCardCriteriaType
     */
    cardDetails?: PrepaidCardDetailsType;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof PrepaidCardCriteriaType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof PrepaidCardCriteriaType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof PrepaidCardCriteriaType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof PrepaidCardCriteriaType
     */
    purgeDate?: Date;
    /**
     * Indicates a gift card type.
     * @type {boolean}
     * @memberof PrepaidCardCriteriaType
     */
    giftCard?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PrepaidCardCriteriaType
     */
    postChargeToRoom?: boolean;
    /**
     *
     * @type {SaleCriteriaType}
     * @memberof PrepaidCardCriteriaType
     */
    saleCriteria?: SaleCriteriaType;
    /**
     *
     * @type {number}
     * @memberof PrepaidCardCriteriaType
     */
    vendorInterfaceID?: number;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof PrepaidCardCriteriaType
     */
    cashierId?: number;
}
/**
 * Check if a given object implements the PrepaidCardCriteriaType interface.
 */
export declare function instanceOfPrepaidCardCriteriaType(value: object): boolean;
export declare function PrepaidCardCriteriaTypeFromJSON(json: any): PrepaidCardCriteriaType;
export declare function PrepaidCardCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PrepaidCardCriteriaType;
export declare function PrepaidCardCriteriaTypeToJSON(value?: PrepaidCardCriteriaType | null): any;
