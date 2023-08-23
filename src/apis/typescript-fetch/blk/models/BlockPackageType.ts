/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
import type { PackageCodeHeaderType } from './PackageCodeHeaderType';
import {
    PackageCodeHeaderTypeFromJSON,
    PackageCodeHeaderTypeFromJSONTyped,
    PackageCodeHeaderTypeToJSON,
} from './PackageCodeHeaderType';
import type { PackageConsumptionType } from './PackageConsumptionType';
import {
    PackageConsumptionTypeFromJSON,
    PackageConsumptionTypeFromJSONTyped,
    PackageConsumptionTypeToJSON,
} from './PackageConsumptionType';
import type { TimeSpanType } from './TimeSpanType';
import {
    TimeSpanTypeFromJSON,
    TimeSpanTypeFromJSONTyped,
    TimeSpanTypeToJSON,
} from './TimeSpanType';

/**
 * This contains the full information and schedule of the block package.
 * @export
 * @interface BlockPackageType
 */
export interface BlockPackageType {
    /**
     * 
     * @type {PackageCodeHeaderType}
     * @memberof BlockPackageType
     */
    packageHeaderType?: PackageCodeHeaderType;
    /**
     * 
     * @type {TimeSpanType}
     * @memberof BlockPackageType
     */
    newTimeSpan?: TimeSpanType;
    /**
     * 
     * @type {PackageConsumptionType}
     * @memberof BlockPackageType
     */
    consumptionDetails?: PackageConsumptionType;
    /**
     * Package code. This is the unique code used for the package and is a required element.
     * @type {string}
     * @memberof BlockPackageType
     */
    packageCode?: string;
    /**
     * The rate code which contains this package. If the package is not part of a rate code, this will be empty. Required element and part of the key to fetch the correct package record.
     * @type {string}
     * @memberof BlockPackageType
     */
    ratePlanCode?: string;
    /**
     * Required value when changing a package. If the original start date was null, then null is required.
     * @type {Date}
     * @memberof BlockPackageType
     */
    startDate?: Date;
    /**
     * Required value when changing a package. If the original end date was null, then null is required.
     * @type {Date}
     * @memberof BlockPackageType
     */
    endDate?: Date;
    /**
     * Package group code. If this package is part of a package group, the group code is indicated here. This is a required element and is part of the key to fetch the correct package record .
     * @type {string}
     * @memberof BlockPackageType
     */
    packageGroup?: string;
}

/**
 * Check if a given object implements the BlockPackageType interface.
 */
export function instanceOfBlockPackageType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockPackageTypeFromJSON(json: any): BlockPackageType {
    return BlockPackageTypeFromJSONTyped(json, false);
}

export function BlockPackageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockPackageType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'packageHeaderType': !exists(json, 'packageHeaderType') ? undefined : PackageCodeHeaderTypeFromJSON(json['packageHeaderType']),
        'newTimeSpan': !exists(json, 'newTimeSpan') ? undefined : TimeSpanTypeFromJSON(json['newTimeSpan']),
        'consumptionDetails': !exists(json, 'consumptionDetails') ? undefined : PackageConsumptionTypeFromJSON(json['consumptionDetails']),
        'packageCode': !exists(json, 'packageCode') ? undefined : json['packageCode'],
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !exists(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'packageGroup': !exists(json, 'packageGroup') ? undefined : json['packageGroup'],
    };
}

export function BlockPackageTypeToJSON(value?: BlockPackageType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'packageHeaderType': PackageCodeHeaderTypeToJSON(value.packageHeaderType),
        'newTimeSpan': TimeSpanTypeToJSON(value.newTimeSpan),
        'consumptionDetails': PackageConsumptionTypeToJSON(value.consumptionDetails),
        'packageCode': value.packageCode,
        'ratePlanCode': value.ratePlanCode,
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0,10)),
        'packageGroup': value.packageGroup,
    };
}

