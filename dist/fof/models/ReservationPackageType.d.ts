/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PackageCodeHeaderType } from './PackageCodeHeaderType';
import type { PackageConsumptionType } from './PackageConsumptionType';
import type { ProductSourceType } from './ProductSourceType';
import type { ReservationPackageScheduleType } from './ReservationPackageScheduleType';
import type { TimeSpanType } from './TimeSpanType';
/**
 * A ReservationPackageType class.
 * @export
 * @interface ReservationPackageType
 */
export interface ReservationPackageType {
    /**
     *
     * @type {PackageCodeHeaderType}
     * @memberof ReservationPackageType
     */
    packageHeaderType?: PackageCodeHeaderType;
    /**
     * A HotelPackageSchedule type.
     * @type {Array<ReservationPackageScheduleType>}
     * @memberof ReservationPackageType
     */
    scheduleList?: Array<ReservationPackageScheduleType>;
    /**
     *
     * @type {TimeSpanType}
     * @memberof ReservationPackageType
     */
    newTimeSpan?: TimeSpanType;
    /**
     *
     * @type {PackageConsumptionType}
     * @memberof ReservationPackageType
     */
    consumptionDetails?: PackageConsumptionType;
    /**
     * Package code. This is the unique code used for the package and is a required element.
     * @type {string}
     * @memberof ReservationPackageType
     */
    packageCode?: string;
    /**
     * Reservation Package Opera Internal Unique Id. This is the unique Id used for this reservation package.
     * @type {number}
     * @memberof ReservationPackageType
     */
    internalID?: number;
    /**
     * The rate code which contains this package. If the package is not part of a rate code, this will be empty. Required element and part of the key to fetch the correct package record on the reservation.
     * @type {string}
     * @memberof ReservationPackageType
     */
    ratePlanCode?: string;
    /**
     * Required value when changing a reservation package. If the original start date was null, then null is required.
     * @type {Date}
     * @memberof ReservationPackageType
     */
    startDate?: Date;
    /**
     * Required value when changing a reservation package. If the original end date was null, then null is required.
     * @type {Date}
     * @memberof ReservationPackageType
     */
    endDate?: Date;
    /**
     * Package group code. If this package is part of a package group, the group code is indicated here. This is a required element and is part of the key to fetch the correct package record .
     * @type {string}
     * @memberof ReservationPackageType
     */
    packageGroup?: string;
    /**
     *
     * @type {ProductSourceType}
     * @memberof ReservationPackageType
     */
    source?: ProductSourceType;
    /**
     * This is the Award code used to redeem the package if the package is a redemption package.
     * @type {string}
     * @memberof ReservationPackageType
     */
    awardCode?: string;
    /**
     * Indicates the points used to redeem the redemption package.
     * @type {number}
     * @memberof ReservationPackageType
     */
    points?: number;
}
/**
 * Check if a given object implements the ReservationPackageType interface.
 */
export declare function instanceOfReservationPackageType(value: object): boolean;
export declare function ReservationPackageTypeFromJSON(json: any): ReservationPackageType;
export declare function ReservationPackageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationPackageType;
export declare function ReservationPackageTypeToJSON(value?: ReservationPackageType | null): any;
