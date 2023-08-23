/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { GuestHousekeepingServiceRequestType } from './GuestHousekeepingServiceRequestType';
import {
    GuestHousekeepingServiceRequestTypeFromJSON,
    GuestHousekeepingServiceRequestTypeFromJSONTyped,
    GuestHousekeepingServiceRequestTypeToJSON,
} from './GuestHousekeepingServiceRequestType';
import type { PointsType } from './PointsType';
import {
    PointsTypeFromJSON,
    PointsTypeFromJSONTyped,
    PointsTypeToJSON,
} from './PointsType';
import type { PromotionType } from './PromotionType';
import {
    PromotionTypeFromJSON,
    PromotionTypeFromJSONTyped,
    PromotionTypeToJSON,
} from './PromotionType';
import type { ResExpectedTimesType } from './ResExpectedTimesType';
import {
    ResExpectedTimesTypeFromJSON,
    ResExpectedTimesTypeFromJSONTyped,
    ResExpectedTimesTypeToJSON,
} from './ResExpectedTimesType';
import type { ResGuaranteeType } from './ResGuaranteeType';
import {
    ResGuaranteeTypeFromJSON,
    ResGuaranteeTypeFromJSONTyped,
    ResGuaranteeTypeToJSON,
} from './ResGuaranteeType';
import type { ReservationBlockType } from './ReservationBlockType';
import {
    ReservationBlockTypeFromJSON,
    ReservationBlockTypeFromJSONTyped,
    ReservationBlockTypeToJSON,
} from './ReservationBlockType';
import type { TimeSpanType } from './TimeSpanType';
import {
    TimeSpanTypeFromJSON,
    TimeSpanTypeFromJSONTyped,
    TimeSpanTypeToJSON,
} from './TimeSpanType';

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
     * True indicates room type is a component type.
     * @type {boolean}
     * @memberof StayInfoType
     */
    componentRoomType?: boolean;
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
    roomId?: string;
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
    linkCode?: string;
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
     * Description of Market code
     * @type {string}
     * @memberof StayInfoType
     */
    marketDescription?: string;
    /**
     * Source of business
     * @type {string}
     * @memberof StayInfoType
     */
    sourceCode?: string;
    /**
     * Description of the source of business.
     * @type {string}
     * @memberof StayInfoType
     */
    sourceCodeDescription?: string;
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
    /**
     * Represents the room was assigned by AI Room Assignment.
     * @type {boolean}
     * @memberof StayInfoType
     */
    assignedByAI?: boolean;
    /**
     * Represents the room was assigned by AI Room Assignment.
     * @type {boolean}
     * @memberof StayInfoType
     */
    upgradedByAI?: boolean;
}

/**
 * Check if a given object implements the StayInfoType interface.
 */
export function instanceOfStayInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StayInfoTypeFromJSON(json: any): StayInfoType {
    return StayInfoTypeFromJSONTyped(json, false);
}

export function StayInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StayInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'arrivalDate': !exists(json, 'arrivalDate') ? undefined : (new Date(json['arrivalDate'])),
        'departureDate': !exists(json, 'departureDate') ? undefined : (new Date(json['departureDate'])),
        'originalTimeSpan': !exists(json, 'originalTimeSpan') ? undefined : TimeSpanTypeFromJSON(json['originalTimeSpan']),
        'expectedTimes': !exists(json, 'expectedTimes') ? undefined : ResExpectedTimesTypeFromJSON(json['expectedTimes']),
        'adultCount': !exists(json, 'adultCount') ? undefined : json['adultCount'],
        'childCount': !exists(json, 'childCount') ? undefined : json['childCount'],
        'roomClass': !exists(json, 'roomClass') ? undefined : json['roomClass'],
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
        'componentRoomType': !exists(json, 'componentRoomType') ? undefined : json['componentRoomType'],
        'numberOfRooms': !exists(json, 'numberOfRooms') ? undefined : json['numberOfRooms'],
        'roomId': !exists(json, 'roomId') ? undefined : json['roomId'],
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'rateAmount': !exists(json, 'rateAmount') ? undefined : CurrencyAmountTypeFromJSON(json['rateAmount']),
        'points': !exists(json, 'points') ? undefined : PointsTypeFromJSON(json['points']),
        'rateSuppressed': !exists(json, 'rateSuppressed') ? undefined : json['rateSuppressed'],
        'reservationBlock': !exists(json, 'reservationBlock') ? undefined : ReservationBlockTypeFromJSON(json['reservationBlock']),
        'bookingChannelCode': !exists(json, 'bookingChannelCode') ? undefined : json['bookingChannelCode'],
        'linkCode': !exists(json, 'linkCode') ? undefined : json['linkCode'],
        'fixedRate': !exists(json, 'fixedRate') ? undefined : json['fixedRate'],
        'totalAmount': !exists(json, 'totalAmount') ? undefined : CurrencyAmountTypeFromJSON(json['totalAmount']),
        'guarantee': !exists(json, 'guarantee') ? undefined : ResGuaranteeTypeFromJSON(json['guarantee']),
        'promotion': !exists(json, 'promotion') ? undefined : PromotionTypeFromJSON(json['promotion']),
        'marketCode': !exists(json, 'marketCode') ? undefined : json['marketCode'],
        'marketDescription': !exists(json, 'marketDescription') ? undefined : json['marketDescription'],
        'sourceCode': !exists(json, 'sourceCode') ? undefined : json['sourceCode'],
        'sourceCodeDescription': !exists(json, 'sourceCodeDescription') ? undefined : json['sourceCodeDescription'],
        'balance': !exists(json, 'balance') ? undefined : CurrencyAmountTypeFromJSON(json['balance']),
        'compBalance': !exists(json, 'compBalance') ? undefined : CurrencyAmountTypeFromJSON(json['compBalance']),
        'roomTypeCharged': !exists(json, 'roomTypeCharged') ? undefined : json['roomTypeCharged'],
        'depositPayments': !exists(json, 'depositPayments') ? undefined : CurrencyAmountTypeFromJSON(json['depositPayments']),
        'guestServiceStatus': !exists(json, 'guestServiceStatus') ? undefined : GuestHousekeepingServiceRequestTypeFromJSON(json['guestServiceStatus']),
        'scheduledCheckoutTime': !exists(json, 'scheduledCheckoutTime') ? undefined : (new Date(json['scheduledCheckoutTime'])),
        'roomNumberLocked': !exists(json, 'roomNumberLocked') ? undefined : json['roomNumberLocked'],
        'pseudoRoom': !exists(json, 'pseudoRoom') ? undefined : json['pseudoRoom'],
        'assignedByAI': !exists(json, 'assignedByAI') ? undefined : json['assignedByAI'],
        'upgradedByAI': !exists(json, 'upgradedByAI') ? undefined : json['upgradedByAI'],
    };
}

export function StayInfoTypeToJSON(value?: StayInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'arrivalDate': value.arrivalDate === undefined ? undefined : (value.arrivalDate.toISOString().substr(0,10)),
        'departureDate': value.departureDate === undefined ? undefined : (value.departureDate.toISOString().substr(0,10)),
        'originalTimeSpan': TimeSpanTypeToJSON(value.originalTimeSpan),
        'expectedTimes': ResExpectedTimesTypeToJSON(value.expectedTimes),
        'adultCount': value.adultCount,
        'childCount': value.childCount,
        'roomClass': value.roomClass,
        'roomType': value.roomType,
        'componentRoomType': value.componentRoomType,
        'numberOfRooms': value.numberOfRooms,
        'roomId': value.roomId,
        'ratePlanCode': value.ratePlanCode,
        'rateAmount': CurrencyAmountTypeToJSON(value.rateAmount),
        'points': PointsTypeToJSON(value.points),
        'rateSuppressed': value.rateSuppressed,
        'reservationBlock': ReservationBlockTypeToJSON(value.reservationBlock),
        'bookingChannelCode': value.bookingChannelCode,
        'linkCode': value.linkCode,
        'fixedRate': value.fixedRate,
        'totalAmount': CurrencyAmountTypeToJSON(value.totalAmount),
        'guarantee': ResGuaranteeTypeToJSON(value.guarantee),
        'promotion': PromotionTypeToJSON(value.promotion),
        'marketCode': value.marketCode,
        'marketDescription': value.marketDescription,
        'sourceCode': value.sourceCode,
        'sourceCodeDescription': value.sourceCodeDescription,
        'balance': CurrencyAmountTypeToJSON(value.balance),
        'compBalance': CurrencyAmountTypeToJSON(value.compBalance),
        'roomTypeCharged': value.roomTypeCharged,
        'depositPayments': CurrencyAmountTypeToJSON(value.depositPayments),
        'guestServiceStatus': GuestHousekeepingServiceRequestTypeToJSON(value.guestServiceStatus),
        'scheduledCheckoutTime': value.scheduledCheckoutTime === undefined ? undefined : (value.scheduledCheckoutTime.toISOString().substr(0,10)),
        'roomNumberLocked': value.roomNumberLocked,
        'pseudoRoom': value.pseudoRoom,
        'assignedByAI': value.assignedByAI,
        'upgradedByAI': value.upgradedByAI,
    };
}

