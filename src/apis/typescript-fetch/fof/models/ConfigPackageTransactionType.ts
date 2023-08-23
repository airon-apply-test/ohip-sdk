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
import type { PackagePostingRulesType } from './PackagePostingRulesType';
import {
    PackagePostingRulesTypeFromJSON,
    PackagePostingRulesTypeFromJSONTyped,
    PackagePostingRulesTypeToJSON,
} from './PackagePostingRulesType';

/**
 * A HotelPackageTransaction type.
 * @export
 * @interface ConfigPackageTransactionType
 */
export interface ConfigPackageTransactionType {
    /**
     * Package is marked as an allowance, in case charge is expected back to the guest account from external interface eg. POS. which need to be offset against a consumption allowance.
     * @type {boolean}
     * @memberof ConfigPackageTransactionType
     */
    allowance?: boolean;
    /**
     * The currency code for this package.
     * @type {string}
     * @memberof ConfigPackageTransactionType
     */
    currency?: string;
    /**
     * The posting frequency for this package, e.g., daily, arrival, departure, etc.
     * @type {string}
     * @memberof ConfigPackageTransactionType
     */
    postingType?: string;
    /**
     * The price calculation rule for this package.
     * @type {string}
     * @memberof ConfigPackageTransactionType
     */
    calculationRule?: string;
    /**
     * 
     * @type {PackagePostingRulesType}
     * @memberof ConfigPackageTransactionType
     */
    packagePostingRules?: PackagePostingRulesType;
}

/**
 * Check if a given object implements the ConfigPackageTransactionType interface.
 */
export function instanceOfConfigPackageTransactionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ConfigPackageTransactionTypeFromJSON(json: any): ConfigPackageTransactionType {
    return ConfigPackageTransactionTypeFromJSONTyped(json, false);
}

export function ConfigPackageTransactionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigPackageTransactionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'allowance': !exists(json, 'allowance') ? undefined : json['allowance'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'postingType': !exists(json, 'postingType') ? undefined : json['postingType'],
        'calculationRule': !exists(json, 'calculationRule') ? undefined : json['calculationRule'],
        'packagePostingRules': !exists(json, 'packagePostingRules') ? undefined : PackagePostingRulesTypeFromJSON(json['packagePostingRules']),
    };
}

export function ConfigPackageTransactionTypeToJSON(value?: ConfigPackageTransactionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'allowance': value.allowance,
        'currency': value.currency,
        'postingType': value.postingType,
        'calculationRule': value.calculationRule,
        'packagePostingRules': PackagePostingRulesTypeToJSON(value.packagePostingRules),
    };
}

