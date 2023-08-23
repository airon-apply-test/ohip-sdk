/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
import type { UniqueIDType } from './UniqueIDType';
/**
 * Holds fixed charge information.
 * @export
 * @interface PrepaidCardType
 */
export interface PrepaidCardType {
    /**
     * Hotel context for the Reservations.
     * @type {string}
     * @memberof PrepaidCardType
     */
    hotelId?: string;
    /**
     *
     * @type {ReservationId}
     * @memberof PrepaidCardType
     */
    reservationId?: ReservationId;
    /**
     * Family name, last name or Company Name.
     * @type {string}
     * @memberof PrepaidCardType
     */
    name?: string;
    /**
     * Given name, first name or names.
     * @type {string}
     * @memberof PrepaidCardType
     */
    firstName?: string;
    /**
     *
     * @type {ProfileId}
     * @memberof PrepaidCardType
     */
    profileId?: ProfileId;
    /**
     * Prepaid card / account number.
     * @type {string}
     * @memberof PrepaidCardType
     */
    cardNo?: string;
    /**
     * Masked Prepaid card / account number.
     * @type {string}
     * @memberof PrepaidCardType
     */
    cardNumberMasked?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardType
     */
    amount?: CurrencyAmountType;
    /**
     * Prepaid card pin code.
     * @type {string}
     * @memberof PrepaidCardType
     */
    pinCode?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof PrepaidCardType
     */
    interfaceId?: UniqueIDType;
    /**
     *
     * @type {PrepaidCardDetailsType}
     * @memberof PrepaidCardType
     */
    cardDetails?: PrepaidCardDetailsType;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof PrepaidCardType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof PrepaidCardType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof PrepaidCardType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof PrepaidCardType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof PrepaidCardType
     */
    purgeDate?: Date;
    /**
     * Indicates a gift card type.
     * @type {boolean}
     * @memberof PrepaidCardType
     */
    giftCard?: boolean;
}
/**
 * Check if a given object implements the PrepaidCardType interface.
 */
export declare function instanceOfPrepaidCardType(value: object): boolean;
export declare function PrepaidCardTypeFromJSON(json: any): PrepaidCardType;
export declare function PrepaidCardTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PrepaidCardType;
export declare function PrepaidCardTypeToJSON(value?: PrepaidCardType | null): any;
