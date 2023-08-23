/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CompTypeType } from './CompTypeType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for creating CompTypes.
 * @export
 * @interface CompTypesCriteria
 */
export interface CompTypesCriteria {
    /**
     * List of Comp Types.
     * @type {Array<CompTypeType>}
     * @memberof CompTypesCriteria
     */
    compTypes?: Array<CompTypeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CompTypesCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CompTypesCriteria
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CompTypesCriteria interface.
 */
export declare function instanceOfCompTypesCriteria(value: object): boolean;
export declare function CompTypesCriteriaFromJSON(json: any): CompTypesCriteria;
export declare function CompTypesCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompTypesCriteria;
export declare function CompTypesCriteriaToJSON(value?: CompTypesCriteria | null): any;
