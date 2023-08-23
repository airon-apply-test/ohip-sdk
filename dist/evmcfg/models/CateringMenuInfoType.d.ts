/**
 * OPERA Cloud Event Configuration API
 * This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { DateRangeType } from './DateRangeType';
import type { MenuTypeType } from './MenuTypeType';
import type { ServingType } from './ServingType';
import type { TranslationTextType2000 } from './TranslationTextType2000';
/**
 *
 * @export
 * @interface CateringMenuInfoType
 */
export interface CateringMenuInfoType {
    /**
     * This type holds name of Menu Class Name.
     * @type {string}
     * @memberof CateringMenuInfoType
     */
    className?: string;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof CateringMenuInfoType
     */
    name?: TranslationTextType2000;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof CateringMenuInfoType
     */
    restriction?: TranslationTextType2000;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof CateringMenuInfoType
     */
    description?: TranslationTextType2000;
    /**
     *
     * @type {Array<string>}
     * @memberof CateringMenuInfoType
     */
    dietaryList?: Array<string>;
    /**
     *
     * @type {Array<string>}
     * @memberof CateringMenuInfoType
     */
    eventTypes?: Array<string>;
    /**
     *
     * @type {MenuTypeType}
     * @memberof CateringMenuInfoType
     */
    type?: MenuTypeType;
    /**
     * This type holds quick insert value, v5 functionality as Article Number for the Menu.
     * @type {string}
     * @memberof CateringMenuInfoType
     */
    quickInsertCode?: string;
    /**
     *
     * @type {ServingType}
     * @memberof CateringMenuInfoType
     */
    servingType?: ServingType;
    /**
     * This type holds TableCapacity for the given Menu.
     * @type {string}
     * @memberof CateringMenuInfoType
     */
    servingSize?: string;
    /**
     * Return true, when all Menu Items added will be marked as Included and there will be a Global Price for this menu
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    includedInMenu?: boolean;
    /**
     * This type holds value of consumption, Menu items will be charged on a consumption basis.
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    consumptionBased?: boolean;
    /**
     *
     * @type {DateRangeType}
     * @memberof CateringMenuInfoType
     */
    sellDateRange?: DateRangeType;
    /**
     *
     * @type {DateRangeType}
     * @memberof CateringMenuInfoType
     */
    eventDateRange?: DateRangeType;
    /**
     * The total number of menus sold for the day on which the event takes place.
     * @type {number}
     * @memberof CateringMenuInfoType
     */
    menusSold?: number;
    /**
     * This type tells about whether menu is book through web or not.
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    webBookable?: boolean;
    /**
     * Selection will note this Menu is Inactive and unable to be added to a Catering Event.
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    inactive?: boolean;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CateringMenuInfoType
     */
    salesPrice?: CurrencyAmountType;
    /**
     * Return true means, menu item as Included in the Menu Price
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    includedInPrice?: boolean;
    /**
     * Indicator of this being a multi-choice menu
     * @type {boolean}
     * @memberof CateringMenuInfoType
     */
    multiChoice?: boolean;
    /**
     * Define the Courses populated into the Multi Choice Menu Configuration
     * @type {number}
     * @memberof CateringMenuInfoType
     */
    courseCount?: number;
}
/**
 * Check if a given object implements the CateringMenuInfoType interface.
 */
export declare function instanceOfCateringMenuInfoType(value: object): boolean;
export declare function CateringMenuInfoTypeFromJSON(json: any): CateringMenuInfoType;
export declare function CateringMenuInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CateringMenuInfoType;
export declare function CateringMenuInfoTypeToJSON(value?: CateringMenuInfoType | null): any;
