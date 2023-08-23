/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
/**
 * Package Consumption Information. Includes information on quantity of the package used, allowance consumption etc.
 * @export
 * @interface PackageConsumptionType
 */
export interface PackageConsumptionType {
    /**
     * The default quantity attached in the package.
     * @type {number}
     * @memberof PackageConsumptionType
     */
    defaultQuantity?: number;
    /**
     * The quantity which has been excluded in the package.
     * @type {number}
     * @memberof PackageConsumptionType
     */
    excludedQuantity?: number;
    /**
     * The total quantity of the package, calculated based on the calculation rule as defined in the PackageHeaderType.
     * @type {number}
     * @memberof PackageConsumptionType
     */
    totalQuantity?: number;
    /**
     * Indicates if Allowance(for POS packages) has been consumed/posted for today.
     * @type {boolean}
     * @memberof PackageConsumptionType
     */
    allowanceConsumed?: boolean;
}

/**
 * Check if a given object implements the PackageConsumptionType interface.
 */
export function instanceOfPackageConsumptionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PackageConsumptionTypeFromJSON(json: any): PackageConsumptionType {
    return PackageConsumptionTypeFromJSONTyped(json, false);
}

export function PackageConsumptionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageConsumptionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'defaultQuantity': !exists(json, 'defaultQuantity') ? undefined : json['defaultQuantity'],
        'excludedQuantity': !exists(json, 'excludedQuantity') ? undefined : json['excludedQuantity'],
        'totalQuantity': !exists(json, 'totalQuantity') ? undefined : json['totalQuantity'],
        'allowanceConsumed': !exists(json, 'allowanceConsumed') ? undefined : json['allowanceConsumed'],
    };
}

export function PackageConsumptionTypeToJSON(value?: PackageConsumptionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'defaultQuantity': value.defaultQuantity,
        'excludedQuantity': value.excludedQuantity,
        'totalQuantity': value.totalQuantity,
        'allowanceConsumed': value.allowanceConsumed,
    };
}

