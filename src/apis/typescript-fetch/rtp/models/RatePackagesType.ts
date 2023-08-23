/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RatePackageCodeType } from './RatePackageCodeType';
import {
    RatePackageCodeTypeFromJSON,
    RatePackageCodeTypeFromJSONTyped,
    RatePackageCodeTypeToJSON,
} from './RatePackageCodeType';
import type { RatePackageGroupType } from './RatePackageGroupType';
import {
    RatePackageGroupTypeFromJSON,
    RatePackageGroupTypeFromJSONTyped,
    RatePackageGroupTypeToJSON,
} from './RatePackageGroupType';

/**
 * 
 * @export
 * @interface RatePackagesType
 */
export interface RatePackagesType {
    /**
     * Rate Package Full Information
     * @type {Array<RatePackageCodeType>}
     * @memberof RatePackagesType
     */
    packages?: Array<RatePackageCodeType>;
    /**
     * Package Group full Information along with the members that belong to this Group.
     * @type {Array<RatePackageGroupType>}
     * @memberof RatePackagesType
     */
    packageGroups?: Array<RatePackageGroupType>;
}

/**
 * Check if a given object implements the RatePackagesType interface.
 */
export function instanceOfRatePackagesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RatePackagesTypeFromJSON(json: any): RatePackagesType {
    return RatePackagesTypeFromJSONTyped(json, false);
}

export function RatePackagesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RatePackagesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'packages': !exists(json, 'packages') ? undefined : ((json['packages'] as Array<any>).map(RatePackageCodeTypeFromJSON)),
        'packageGroups': !exists(json, 'packageGroups') ? undefined : ((json['packageGroups'] as Array<any>).map(RatePackageGroupTypeFromJSON)),
    };
}

export function RatePackagesTypeToJSON(value?: RatePackagesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'packages': value.packages === undefined ? undefined : ((value.packages as Array<any>).map(RatePackageCodeTypeToJSON)),
        'packageGroups': value.packageGroups === undefined ? undefined : ((value.packageGroups as Array<any>).map(RatePackageGroupTypeToJSON)),
    };
}

