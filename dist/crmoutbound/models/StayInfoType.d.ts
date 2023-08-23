/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { GuestHousekeepingServiceRequestType } from './GuestHousekeepingServiceRequestType';
import type { PointsType } from './PointsType';
import type { PromotionType } from './PromotionType';
import type { ResExpectedTimesType } from './ResExpectedTimesType';
import type { ResGuaranteeType } from './ResGuaranteeType';
import type { ReservationBlockType } from './ReservationBlockType';
import type { TimeSpanType } from './TimeSpanType';
/**
 * Details on the Room Stay including Guest Counts, Time Span of this Room Stay, pointers to Res Guests, guest Memberships, Comments and Special Requests pertaining to this particular Room Stay and finally financial information related to the Room Stay, including Guarantee, Deposit and Payment and Cancellation Penalties.
 * @export
 * @interface StayInfoType
 */
export interface StayInfoType {
    /**
     *
     * @type {Date}
     * @memberof StayInfoType
     */
    arrivalDate?: Date;
    /**
     *
     * @type {Date}
     * @memberof StayInfoType
     */
    departureDate?: Date;
    /**
     *
     * @type {TimeSpanType}
     * @memberof StayInfoType
     */
    originalTimeSpan?: TimeSpanType;
    /**
     *
     * @type {ResExpectedTimesType}
     * @memberof StayInfoType
     */
    expectedTimes?: ResExpectedTimesType;
    /**
     * A collection of Guest Counts associated with Room Stay.
     * @type {number}
     * @memberof StayInfoType
     */
    adultCount?: number;
    /**
     * A collection of Child Counts associated with Room Stay.
     * @type {number}
     * @memberof StayInfoType
     */
    childCount?: number;
    /**
     * Room class code
     * @type {string}
     * @memberof StayInfoType
     */
    roomClass?: string;
    /**
     * Room type code
     * @type {string}
     * @memberof StayInfoType
     */
    roomType?: string;
    /**
     * Room Id
     * @type {number}
     * @memberof StayInfoType
     */
    numberOfRooms?: number;
    /**
     * Room Id
     * @type {string}
     * @memberof StayInfoType
     */
    roomNumber?: string;
    /**
     * Rate plan code
     * @type {string}
     * @memberof StayInfoType
     */
    ratePlanCode?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof StayInfoType
     */
    rateAmount?: CurrencyAmountType;
    /**
     *
     * @type {PointsType}
     * @memberof StayInfoType
     */
    points?: PointsType;
    /**
     * Whether this rate should be suppressed from view
     * @type {boolean}
     * @memberof StayInfoType
     */
    rateSuppressed?: boolean;
    /**
     *
     * @type {ReservationBlockType}
     * @memberof StayInfoType
     */
    reservationBlock?: ReservationBlockType;
    /**
     * Booking channel code
     * @type {string}
     * @memberof StayInfoType
     */
    bookingChannelCode?: string;
    /**
     * Party code
     * @type {string}
     * @memberof StayInfoType
     */
    partyCode?: string;
    /**
     * True if the rate is a fixed rate, otherwise false
     * @type {boolean}
     * @memberof StayInfoType
     */
    fixedRate?: boolean;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof StayInfoType
     */
    totalAmount?: CurrencyAmountType;
    /**
     *
     * @type {ResGuaranteeType}
     * @memberof StayInfoType
     */
    guarantee?: ResGuaranteeType;
    /**
     *
     * @type {PromotionType}
     * @memberof StayInfoType
     */
    promotion?: PromotionType;
    /**
     * Market code
     * @type {string}
     * @memberof StayInfoType
     */
    marketCode?: string;
    /**
     * Source of business
     * @type {string}
     * @memberof StayInfoType
     */
    sourceOfBusiness?: string;
    /**
     * Description of the source of business.
     * @type {string}
     * @memberof StayInfoType
     */
    sourceOfBusinessDescription?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof StayInfoType
     */
    balance?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof StayInfoType
     */
    compBalance?: CurrencyAmountType;
    /**
     * Room type code that was charged
     * @type {string}
     * @memberof StayInfoType
     */
    roomTypeCharged?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof StayInfoType
     */
    depositPayments?: CurrencyAmountType;
    /**
     *
     * @type {GuestHousekeepingServiceRequestType}
     * @memberof StayInfoType
     */
    guestServiceStatus?: GuestHousekeepingServiceRequestType;
    /**
     * Indicates that this reservation is scheduled for automated check out.
     * @type {Date}
     * @memberof StayInfoType
     */
    scheduledCheckoutTime?: Date;
    /**
     * When true, indicates a room number cannot be changed. When false, indicates a room number may be changed.
     * @type {boolean}
     * @memberof StayInfoType
     */
    roomNumberLocked?: boolean;
    /**
     * True indicates as pseudo room type. This is usually used for a posting master reservation.
     * @type {boolean}
     * @memberof StayInfoType
     */
    pseudoRoom?: boolean;
}
/**
 * Check if a given object implements the StayInfoType interface.
 */
export declare function instanceOfStayInfoType(value: object): boolean;
export declare function StayInfoTypeFromJSON(json: any): StayInfoType;
export declare function StayInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StayInfoType;
export declare function StayInfoTypeToJSON(value?: StayInfoType | null): any;
